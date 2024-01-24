package entry

import (
	v1 "api-gateway/internal/api/v1"
	custom_middleware "api-gateway/internal/domain/middleware"
	"api-gateway/internal/utils"

	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouterHandler(router *echo.Echo, serverSocket *socketio.Server) {
	router.Use(middleware.Recover())

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))
	router.Use(middleware.BodyLimit("40M"))

	router.Use(custom_middleware.ZapLogger(utils.LoggerZap))

	version1 := router.Group("/api/v1")
	v1.Setup(version1, serverSocket)
}
