package v1

import (
	c "api-gateway/internal/controller/director"
	mid "api-gateway/internal/domain/middleware"
	v "api-gateway/internal/domain/validation"

	"github.com/labstack/echo/v4"
)

func SetDirectorRoutes(router *echo.Group, auth echo.MiddlewareFunc) {
	path := router.Group("/director")

	path.POST("", c.Insert, auth, v.ValidateDirector)
	path.GET("", c.List, mid.ValidatePageQueryParam, mid.ValidatePageSizeQueryParam)
	path.PUT("/:ID", c.Update, auth, v.ValidateDirector, v.ValidateParameterID)
	path.DELETE("/:ID", c.Delete, auth, v.ValidateParameterID)
}
