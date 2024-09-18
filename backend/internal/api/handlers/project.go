package handlers

import (
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Infinities-ICT-Solutions/project-dashboard/data"
	"github.com/Infinities-ICT-Solutions/project-dashboard/internal/auth"
	st "github.com/Infinities-ICT-Solutions/project-dashboard/internal/storage"
	"github.com/Infinities-ICT-Solutions/project-dashboard/internal/utils"
)

type ProjectHandler struct {
	storage st.ProjectStorage
}

func NewProjectHandler(storage st.ProjectStorage) *ProjectHandler {
	return &ProjectHandler{storage: storage}
}

func (p *ProjectHandler) ListProjects(c echo.Context) error {
	userToken := c.Request().Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	roles := userToken.CustomClaims.(*auth.CustomClaims).Roles
	user_id := userToken.RegisteredClaims.Subject

	filter := bson.M{}
	if utils.Contains(roles, "admin") {
		fmt.Println("user is admin")
	} else if utils.Contains(roles, "member") {
		filter = bson.M{
			"$or": bson.A{
				bson.M{"team.projectlead.id": user_id},
				bson.M{"team.teammembers.id": user_id},
			},
		}
		fmt.Println("user is member")
	} else if utils.Contains(roles, "client") {
		filter = bson.M{"clientrepresentative.id": user_id}
		fmt.Println("user is client")
	} else {
		fmt.Println("user is unknown")
		return c.String(http.StatusForbidden, "unauthorized")
	}

	res, err := p.storage.GetProjects(filter)
	// res, err := p.storage.GetProjects(bson.M{})
	if err != nil {
		return c.String(500, "error getting projects")
	}

	return c.JSON(http.StatusOK, res)
}

func (p *ProjectHandler) GetProject(c echo.Context) error {
	id := c.Param("id")
	res, err := p.storage.GetProject(id)
	if err != nil {
		return &echo.HTTPError{Code: 404, Message: "project not found"}
	}

	return c.JSON(http.StatusOK, res)
}

func (p *ProjectHandler) CreateProject(c echo.Context) error {
	// uses echo bind to get form values
	project := &data.Project{}
	c.Bind(project)
	errors, err := project.Validate()
	if err != nil {
		return c.JSON(400, errors)
	}

	// project.StartDate = primitive.NewDateTimeFromTime(time.Now())

	res, err := p.storage.CreateProject(project)
	if err != nil {
		return c.String(500, "error creating project")
	}

	return c.JSON(http.StatusCreated, res)
}

func (p *ProjectHandler) UpdateProject(c echo.Context) error {
	id := c.Param("id")
	project := &data.Project{}
	c.Bind(project)
	errors, err := project.Validate()
	if err != nil {
		return c.JSON(400, errors)
	}

	res, err := p.storage.UpdateProject(id, project)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(http.StatusAccepted, res)
}

func (p *ProjectHandler) UpdateMetrics(c echo.Context) error {
	userToken := c.Request().Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	roles := userToken.CustomClaims.(*auth.CustomClaims).Roles
	user_id := userToken.RegisteredClaims.Subject

	if utils.Contains(roles, "admin") {
		fmt.Println("user is admin")
	} else if utils.Contains(roles, "member") {
		//check if user is project lead
		fmt.Println("user is member")
		project, err := p.storage.GetProject(c.Param("id"))
		if err != nil {
			return &echo.HTTPError{Code: 404, Message: "project not found"}
		}

		if project.Team.ProjectLead.ID != user_id {
			return &echo.HTTPError{Code: 403, Message: "unauthorized"}
		}
	} else {
		fmt.Println("user is unknown")
		return c.String(http.StatusBadRequest, "unknown user role")
	}

	id := c.Param("id")

	project, err := p.storage.GetProject(id)
	if err != nil {
		return &echo.HTTPError{Code: 400, Message: "project not found"}
	}

	if project.Status == "completed" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "can't update a completed project"}
	}

	metrics := &data.Metrics{}
	err = c.Bind(metrics)
	if err != nil {
		return &echo.HTTPError{Code: 400, Message: "invalid metrics"}
	}

	errors, err := metrics.Validate()
	if err != nil {
		return c.JSON(400, errors)
	}

	fmt.Println("metrics", metrics.TotalProjectCostPerMilestone)

	res, err := p.storage.UpdateMetrics(id, metrics)
	if err != nil {
		return &echo.HTTPError{Code: 404, Message: "metrics update failed"}
	}

	return c.JSON(http.StatusOK, res)
}

func (p *ProjectHandler) MarkCompleted(c echo.Context) error {
	userToken := c.Request().Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	roles := userToken.CustomClaims.(*auth.CustomClaims).Roles
	user_id := userToken.RegisteredClaims.Subject

	if utils.Contains(roles, "admin") {
		fmt.Println("user is admin")
	} else if utils.Contains(roles, "member") {
		//check if user is project lead
		fmt.Println("user is member")
		project, err := p.storage.GetProject(c.Param("id"))
		if err != nil {
			return &echo.HTTPError{Code: 404, Message: "project not found"}
		}

		if project.Team.ProjectLead.ID != user_id {
			return &echo.HTTPError{Code: 403, Message: "unauthorized"}
		}
	} else {
		fmt.Println("user is unknown")
		return c.String(http.StatusBadRequest, "unknown user role")
	}

	id := c.Param("id")

	payload := &struct {
		CompletionDate primitive.DateTime `json:"completionDate" form:"completionDate"`
	}{}
	err := c.Bind(payload)
	if err != nil {
		return &echo.HTTPError{Code: 400, Message: "invalid date"}
	}

	fmt.Println("mark complete", payload)

	res, err := p.storage.MarkCompleted(id, payload.CompletionDate)
	if err != nil {
		return &echo.HTTPError{Code: 404, Message: "project not found"}
	}

	return c.JSON(http.StatusOK, res)
}

func (p *ProjectHandler) DeleteProject(c echo.Context) error {
	id := c.Param("id")
	err := p.storage.DeleteProject(id)
	if err != nil {
		return c.String(500, "error deleting project")
	}

	return c.JSON(http.StatusNoContent, nil)
}
