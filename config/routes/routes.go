package routes

import (
	controller "github.com/aviabird/go-echo-seed/app/controllers"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	apiGroup *echo.Group
	ctrl     *controller.Controller
}

func InitiateRoutes(v1 *echo.Group, c *controller.Controller) {
	r := &Routes{apiGroup: v1, ctrl: c}
	r.sessionRoutes()
	r.userRoutes()
}

func (r *Routes) sessionRoutes() {
	session := r.apiGroup.Group("/sessions")
	session.GET("/new", r.ctrl.Session.New)
	session.DELETE("/delete", r.ctrl.Session.DELETE)
}

func (r *Routes) userRoutes() {
	guestUsers := r.apiGroup.Group("/users")
	guestUsers.GET("", r.ctrl.User.UserIndex)
}
