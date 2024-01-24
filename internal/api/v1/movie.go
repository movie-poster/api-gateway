package v1

import (
	c "api-gateway/internal/controller/movie"
	mid "api-gateway/internal/domain/middleware"
	v "api-gateway/internal/domain/validation"

	"github.com/labstack/echo/v4"
)

func SetMovieRoutes(router *echo.Group, auth echo.MiddlewareFunc) {
	path := router.Group("/movie")

	path.POST("", c.Insert, auth, v.ValidateMovie)
	path.PUT("/:ID", c.Update, auth, v.ValidateMovie, v.ValidateParameterID)
	path.GET("", c.List, mid.ValidatePageQueryParam, mid.ValidatePageSizeQueryParam)
	path.DELETE("/:ID", c.Delete, auth, v.ValidateParameterID)
	path.GET("/:ID", c.GetById, v.ValidateParameterID)
}
