package controllers

import (
	"app/app/domain"
	"app/app/interfaces/database"
	"app/app/usecase"
	"net/http"
	"strconv"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

// このdata.Sqlはどこで定義されてるのを読んでるんだっけ？
func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, u)
}

// ContextをginのContextに変えてんでー
func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, user)
}
