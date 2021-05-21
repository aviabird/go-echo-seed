package routes

import (
	controller "github.com/aviabird/go-echo-seed/app/controllers"
	"github.com/labstack/echo/v4"
)

func InitiateRoutes(v1 *echo.Group, c *controller.Controller) {
	guestUsers := v1.Group("/users")
	guestUsers.GET("", c.UserIndex)
}
