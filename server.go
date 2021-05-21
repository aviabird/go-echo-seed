package main

import (
	"os"

	"github.com/aviabird/go-echo-seed/app/controllers"
	"github.com/aviabird/go-echo-seed/app/repo"
	"github.com/aviabird/go-echo-seed/config/db"
	"github.com/aviabird/go-echo-seed/config/router"
	"github.com/aviabird/go-echo-seed/config/routes"

	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "dev"
	}

	godotenv.Load(".env." + env)
}

func main() {
	r := router.New()
	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	repo := repo.InitiateRepo(d)
	ctrls := controllers.InitiateControllers(repo)
	routes.InitiateRoutes(v1, ctrls)

	r.Logger.Fatal(r.Start(os.Getenv("ADDR")))
}
