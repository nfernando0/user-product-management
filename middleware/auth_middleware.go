package middleware

import (
	"net/http"
	jwtToken "products/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusBadRequest, Result{
				Code:    http.StatusBadRequest,
				Data:    nil,
				Message: "unauthorized",
			})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, Result{
				Code:    http.StatusUnauthorized,
				Data:    nil,
				Message: err.Error(),
			})
		}

		c.Set("userLogin", claims)
		return next(c)
	}
}
