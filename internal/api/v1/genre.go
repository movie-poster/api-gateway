package v1

import (
	c "api-gateway/internal/controller/genre"
	mid "api-gateway/internal/domain/middleware"
	v "api-gateway/internal/domain/validation"

	"github.com/labstack/echo/v4"
)

func SetGenreRoutes(router *echo.Group, auth echo.MiddlewareFunc) {
	path := router.Group("/genre")

	path.POST("", c.Insert, auth, v.ValidateDirector)
	path.GET("", c.List, mid.ValidatePageQueryParam, mid.ValidatePageSizeQueryParam)
	path.DELETE("/:ID", c.Delete, auth, v.ValidateParameterID)
}
