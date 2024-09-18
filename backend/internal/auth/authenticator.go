package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/Infinities-ICT-Solutions/project-dashboard/config"
	"github.com/Infinities-ICT-Solutions/project-dashboard/internal/utils"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
)

type Authenticator struct {
	Audience string
	ClientID string
	Domain   string
}

type CustomClaims struct {
	Scope string   `json:"scope"`
	Roles []string `json:"https://ci.com.au/roles"`
	Email string   `json:"https://ci.com.au/email"`
} // Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func New(env config.Env) (*Authenticator, error) {
	return &Authenticator{
		Domain:   env.Auth0.App.Domain,
		ClientID: env.Auth0.App.ClientID,
		Audience: env.Auth0.App.Audience,
	}, nil
}

// func (a *Authenticator) Middleware() func(next http.Handler) http.Handler {
func (a *Authenticator) Middleware() echo.MiddlewareFunc {
	issuerURL, err := url.Parse("https://" + a.Domain + "/")
	fmt.Println("EnsureValidToken", issuerURL)
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{a.Audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return echo.WrapMiddleware(func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	})
	// return func(next http.Handler) http.Handler {
	// 	return middleware.CheckJWT(next)
	// }
}

func (a *Authenticator) HasRoles(permissions []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userToken := c.Request().Context().Value(jwtmiddleware.ContextKey{})

			claims, ok := userToken.(*validator.ValidatedClaims)
			if !ok {
				c.Error(errors.New("could not read JWT claims"))
			}

			roles := claims.CustomClaims.(*CustomClaims).Roles

			for _, permission := range permissions {
				if utils.Contains(roles, permission) {
					return next(c)
				}
			}

			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "You do not have permission to access this resource"})
		}
	}
}
