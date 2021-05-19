package main

import (
	"net/http"
	"os"

	db "github.com/aviabird/go-echo-seed/config"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "dev"
	}

	godotenv.Load(".env." + env)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Pankaj rawat")
	})

	d := db.New()
	db.AutoMigrate(d)

	e.Logger.Fatal(e.Start(os.Getenv("ADDR")))
}
