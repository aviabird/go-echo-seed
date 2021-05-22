package controllers

import (
	"net/http"

	"github.com/aviabird/go-echo-seed/app/repo"
	"github.com/labstack/echo/v4"
)

type SessionController struct {
	repo *repo.Repo
}

func InitiateSessionControllers(repo *repo.Repo) *SessionController {
	return &SessionController{
		repo: repo,
	}
}

func (ctrl *SessionController) New(c echo.Context) error {
	return c.JSON(http.StatusOK, "Login api")
}

func (ctrl *SessionController) DELETE(c echo.Context) error {
	return c.JSON(http.StatusOK, "Logout api")
}
