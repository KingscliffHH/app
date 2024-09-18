package handlers

import (
	"fmt"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"

	"github.com/Infinities-ICT-Solutions/project-dashboard/data"
	st "github.com/Infinities-ICT-Solutions/project-dashboard/internal/storage"
)

type userHandler struct {
	storage  st.UserStorage
	s3Bucket string
}

func NewUserHandler(storage st.UserStorage, s3Bucket string) *userHandler {
	return &userHandler{storage: storage, s3Bucket: s3Bucket}
}

func (u *userHandler) ListUsers(c echo.Context) error {

	res2, err := u.storage.GetAll()
	if err != nil {
		return c.String(500, "error getting users")
	}

	return c.JSON(http.StatusOK, res2)
}

func (u *userHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	res, err := u.storage.GetById(id)
	if err != nil {
		return &echo.HTTPError{Code: 404, Message: "user not found"}
	}

	return c.JSON(http.StatusOK, res)
}

func (u *userHandler) CreateUser(c echo.Context) error {
	// uses echo bind to get form values
	user := &data.User{}
	c.Bind(user)

	errors, err := user.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors)
	}

	res, err := u.storage.Create(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

func (u *userHandler) UploadAvatar(c echo.Context) error {
	file, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	newFilename := fmt.Sprintf("%d", time.Now().UnixNano())

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2"),
	}))
	uploader := s3manager.NewUploader(sess)

	uploadResult, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(u.s3Bucket),
		Key:    aws.String(newFilename),
		Body:   src,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message":  "Upload successful",
		"imageUrl": uploadResult.Location,
	})
}

func (u *userHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := &data.User{}
	c.Bind(user)
	errors, err := user.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors)
	}

	res, err := u.storage.Update(id, user)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(http.StatusAccepted, res)
}

func (u *userHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := u.storage.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (u *userHandler) GetMe(c echo.Context) error {
	userToken := c.Request().Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	res, err := u.storage.GetById(userToken.RegisteredClaims.Subject)
	if err != nil {
		return &echo.HTTPError{Code: 404, Message: "user not found"}
	}

	return c.JSON(http.StatusOK, res)
}
