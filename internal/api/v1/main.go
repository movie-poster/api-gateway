package v1

import (
	"api-gateway/internal/infra/jwt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(router *echo.Group, serverSocket *socketio.Server) {

	jsW := jwt.NewJwtClient()
	auth := middleware.JWTWithConfig(jsW.GetConfig())

	SetActorRoutes(router, auth)
	SetDirectorRoutes(router, auth)
	SetGenreRoutes(router, auth)
	SetMovieRoutes(router, auth)
	SetUserRoutes(router, auth, serverSocket)
}
