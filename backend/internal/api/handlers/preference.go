package handlers

import (
	"fmt"
	"net/http"

	st "github.com/Infinities-ICT-Solutions/project-dashboard/internal/storage"
	"github.com/labstack/echo/v4"
)

type preferenceHander struct {
	storage st.PreferenceStorage
}

func NewPreferenceHandler(storage st.PreferenceStorage) *preferenceHander {
	return &preferenceHander{storage: storage}
}

func (p *preferenceHander) GetOrganisations(c echo.Context) error {
	res, err := p.storage.GetOrganisations()
	if err != nil {
		fmt.Println(err)
		return &echo.HTTPError{Code: 404, Message: "organisations not found"}
	}

	return c.JSON(http.StatusOK, res)
}

// func (p *preferenceHander) AddOrganisation(c echo.Context) error {
// 	id := c.Param("id")
// 	fmt.Println("org", id)
// 	err := p.storage.AddOrganisation(id)
// 	if err != nil {
// 		fmt.Println(err)
// 		return &echo.HTTPError{Code: 404, Message: "can't add org"}
// 	}

// 	return c.JSON(http.StatusOK, nil)
// }
