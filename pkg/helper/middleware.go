package helper

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/wildandhya/mnc-rest-api/internal/model"
)

func ValidateBearer() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authorization := c.Request().Header.Get(echo.HeaderAuthorization)
			if authorization == "" {
				return c.JSON(http.StatusUnauthorized, model.Response{
					Message: "Invalid authorization",
				})
			}

			authValues := strings.Split(authorization, " ")
			authType := strings.ToLower(authValues[0])
			if authType != "bearer" || len(authValues) != 2 {
				return c.JSON(http.StatusUnauthorized, model.Response{
					Message: "Invalid authorization",
				})
			}

			tokenString := authValues[1]
			userId, err := VerifyToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, model.Response{
					Message: "Invalid authorization",
				})
			}
			c.Set("user", userId)
			ctx := context.WithValue(c.Request().Context(), "user", userId)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
