package data

import (
	"fmt"
	"net/mail"
	"strings"
)

type User struct {
	ID           string `json:"id" binding:"required"`
	Type         string `json:"type" form:"type" binding:"required"`
	FullName     string `json:"fullName" form:"fullName" binding:"required"`
	Email        string `json:"email" form:"email" binding:"required"`
	Password     string `json:"password" form:"password" binding:"required"`
	Avatar       string `json:"avatar" form:"avatar" binding:"required"`
	Bio          string `json:"bio" form:"bio" binding:"required"`
	Organisation string `json:"organisation" form:"organisation" binding:"required"`
	ClientRole   string `json:"clientRole" form:"clientRole" binding:"required"`
	LastAccess   string `json:"lastAccess" form:"lastAccess" binding:"required"`
}

func (u *User) Validate() (*ValidationErrorMap, error) {
	errors := make(ValidationErrorMap)

	if strings.TrimSpace(u.Type) == "" {
		errors["type"] = "user type is required"
	}

	if u.Type != "member" && u.Type != "client" {
		errors["type"] = "user type must be either member or client"
	}

	if strings.TrimSpace(u.FullName) == "" {
		errors["fullName"] = "full name is required"
	}

	if strings.TrimSpace(u.Email) == "" {
		errors["email"] = "email is required"
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		errors["email"] = "invalid email"
	}

	if u.Type == "member" && u.Bio == "" {
		errors["bio"] = "bio is required for members"
	}

	if u.Type == "client" && u.Organisation == "" {
		errors["organisation"] = "organisation is required for clients"
	}

	if len(errors) > 0 {
		return &errors, fmt.Errorf("validation error")
	}

	return nil, nil
}
