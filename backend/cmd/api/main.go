package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Infinities-ICT-Solutions/project-dashboard/config"

	// "github.com/Infinities-ICT-Solutions/project-dashboard/internal/api/healthcheck"
	"github.com/Infinities-ICT-Solutions/project-dashboard/internal/api/handlers"
	"github.com/Infinities-ICT-Solutions/project-dashboard/internal/auth"
	"github.com/Infinities-ICT-Solutions/project-dashboard/internal/storage"
	"github.com/Infinities-ICT-Solutions/project-dashboard/pkg/shutdown"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Version = "undefined"

func main() {
	var exitCode int

	defer func() {
		os.Exit(exitCode)
	}()

	// load config
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		exitCode = 1
		return
	}

	// run the server
	cleanup, err := run(env)

	defer cleanup()

	if err != nil {
		fmt.Println("Error running application:", err)
		exitCode = 1
		return
	}

	// ensure graceful shutdown
	shutdown.Gracefully()
}

func run(env config.Env) (func(), error) {
	mongoStorage, err := storage.BootstrapMongo(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	if err != nil {
		return nil, err
	}

	authenticator, err := auth.New(env)
	if err != nil {
		return nil, err
	}

	// create the server
	app := echo.New()

	// middlewares
	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	app.Use(middleware.Logger())

	// app.Static("/assets", "assets")

	// // Index redirects to /projects
	// app.GET("/", func(c echo.Context) error {
	// 	return c.Redirect(http.StatusFound, "/projects")
	// })

	// Users
	pst := storage.NewPreferenceStorage(mongoStorage.DB)
	ust, err := storage.NewUserStorage(env.Auth0.Api.Domain, env.Auth0.Api.ClientID, env.Auth0.Api.ClientSecret, env.Auth0.Api.Audience, pst)

	if err != nil {
		return nil, err
	}

	{
		ph := handlers.NewUserHandler(ust, env.S3Bucket)
		g := app.Group("/users", authenticator.Middleware())
		ga := g.Group("", authenticator.HasRoles([]string{"admin"}))
		ga.POST("", ph.CreateUser)
		ga.GET("/:id", ph.GetUser)
		ga.PUT("/:id", ph.UpdateUser)
		ga.DELETE("/:id", ph.DeleteUser)

		ga.POST("/avatar", ph.UploadAvatar)

		g.GET("", ph.ListUsers)
		g.GET("/me", ph.GetMe)
	}

	// Projects
	{
		st := storage.NewProjectStorage(mongoStorage.DB, ust)
		ph := handlers.NewProjectHandler(st)
		g := app.Group("/projects", authenticator.Middleware())

		g.GET("", ph.ListProjects)
		g.POST("", ph.CreateProject, authenticator.HasRoles([]string{"admin"}))
		g.GET("/:id", ph.GetProject)
		g.PUT("/:id", ph.UpdateProject, authenticator.HasRoles([]string{"admin"}))
		g.DELETE("/:id", ph.DeleteProject, authenticator.HasRoles([]string{"admin"}))

		g.PUT("/:id/metrics", ph.UpdateMetrics, authenticator.HasRoles([]string{"admin", "member"}))
		g.PATCH("/:id/completed", ph.MarkCompleted, authenticator.HasRoles([]string{"admin", "member"}))
	}

	// Benchmarks
	{
		st := storage.NewBenchmarkStorage(mongoStorage.DB)
		ph := handlers.NewBenchmarkHandler(st)
		g := app.Group("/benchmarks", authenticator.Middleware())
		ga := g.Group("", authenticator.HasRoles([]string{"admin"}))

		ga.POST("", ph.CreateBenchmark)
		ga.GET("/:id", ph.GetBenchmark)
		ga.PUT("/:id", ph.UpdateBenchmark)
		ga.DELETE("/:id", ph.DeleteBenchmark)

		g.GET("", ph.ListBenchmarks)
	}

	// Preferences
	{
		st := storage.NewPreferenceStorage(mongoStorage.DB)
		ph := handlers.NewPreferenceHandler(st)
		g := app.Group("/preferences")

		g.GET("/organisations", ph.GetOrganisations)
	}
	// PreferenceStorage
	// Errors
	// app.HTTPErrorHandler = func(err error, c echo.Context) {
	// 	code := http.StatusInternalServerError
	// 	msg := ""
	// 	if he, ok := err.(*echo.HTTPError); ok {
	// 		code = he.Code
	// 		msg = he.Message.(string)
	// 	}
	// 	if code == http.StatusNotFound {
	// 		utils.Render(c, code, errors.NotFound(msg))
	// 		return
	// 	}
	// 	utils.Render(c, code, errors.InternalServerError())
	// }

	// start the server
	go func() {
		err := app.Start(env.ListenAddr)
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()

	return func() {
		app.Close()
	}, nil
}
