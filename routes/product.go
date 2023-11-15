package routes

import (
	"products/db"
	"products/handlers"
	"products/middleware"
	"products/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(db.DB)
	h := handlers.HandlerProduct(productRepository)

	e.GET("/products", h.FindProducts)
	e.GET("/product/:id", h.GetProduct)
	e.POST("/product", middleware.Auth(h.CreateProduct))
	e.PATCH("/product/:id", middleware.Auth(h.UpdateProduct))
	e.DELETE("/product/:id", middleware.Auth(h.DeleteProduct))
}
