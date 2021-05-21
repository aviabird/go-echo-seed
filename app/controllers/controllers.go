package controllers

import "github.com/aviabird/go-echo-seed/app/repo"

type Controller struct {
	repo *repo.Repo
}

func InitiateControllers(repo *repo.Repo) *Controller {
	return &Controller{
		repo: repo,
	}
}
