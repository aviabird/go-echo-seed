package controllers

import "github.com/aviabird/go-echo-seed/app/repo"

type Controller struct {
	User    *UserController
	Session *SessionController
}

func InitiateControllers(repo *repo.Repo) *Controller {
	return &Controller{
		User:    InitiateUserControllers(repo),
		Session: InitiateSessionControllers(repo),
	}
}
