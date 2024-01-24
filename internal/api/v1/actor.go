package v1

import (
	c "api-gateway/internal/controller/actor"
	mid "api-gateway/internal/domain/middleware"
	v "api-gateway/internal/domain/validation"

	"github.com/labstack/echo/v4"
)

func SetActorRoutes(router *echo.Group, auth echo.MiddlewareFunc) {
	path := router.Group("/actor", auth)

	path.POST("", c.Insert, v.ValidateActor)
	path.GET("", c.List, mid.ValidatePageQueryParam, mid.ValidatePageSizeQueryParam)
	path.PUT("/:ID", c.Update, v.ValidateActor, v.ValidateParameterID)
	path.DELETE("/:ID", c.Delete, v.ValidateParameterID)
}
