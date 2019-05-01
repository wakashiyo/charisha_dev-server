package controllers

import (
	"strconv"

	"github.com/wakashiyo/charisha_dev-server/domain"
	"github.com/wakashiyo/charisha_dev-server/interfaces/database"
	"github.com/wakashiyo/charisha_dev-server/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(dbHandler database.DbHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				DbHandler: dbHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	id, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, id)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}
