package controllers

import (
	"net/http"

	"github.com/aviabird/go-echo-seed/app/helpers"
	"github.com/aviabird/go-echo-seed/app/repo"
	"github.com/aviabird/go-echo-seed/app/views"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo *repo.Repo
}

func InitiateUserControllers(repo *repo.Repo) *UserController {
	return &UserController{
		repo: repo,
	}
}

func (ctrl *UserController) UserIndex(c echo.Context) error {
	users, err := ctrl.repo.ListUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.NewError(err))
	}
	if users == nil {
		return c.JSON(http.StatusNotFound, helpers.NotFound())
	}
	return c.JSON(http.StatusOK, views.UserListResponse(users))
}
