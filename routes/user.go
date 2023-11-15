package routes

import (
	"products/db"
	"products/handlers"
	"products/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(db.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUser)
	e.GET("/user/:id", h.GetUser)
	e.PATCH("/user/:id", h.UpdateUser)
}
