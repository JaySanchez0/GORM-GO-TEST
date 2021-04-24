package controller

import (
	"gorm-test/model"
	"gorm-test/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PeopleController struct {
	service *services.PeopleService
}

func (p *PeopleController) GetPeoples(c echo.Context) error {
	return c.JSON(http.StatusAccepted, p.service.GetPeoples())
}

// Add a new people
func (p *PeopleController) PostPeople(c echo.Context) error {
	people := model.People{}
	if c.Bind(&people) != nil {
		return c.String(http.StatusBadRequest, "")
	}
	newPeople, e := p.service.AddPeople(&people)
	if e == nil {
		return c.JSON(http.StatusAccepted, newPeople)
	} else {
		return c.String(http.StatusBadRequest, e.Error())
	}
}

func New(peopleService *services.PeopleService) *PeopleController {
	return &PeopleController{service: peopleService}
}
