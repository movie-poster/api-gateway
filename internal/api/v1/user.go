package v1

import (
	c "api-gateway/internal/controller/user"
	v "api-gateway/internal/domain/validation"

	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
)

func SetUserRoutes(router *echo.Group, auth echo.MiddlewareFunc, serverSocket *socketio.Server) {
	path := router.Group("/user")

	path.POST("", c.Insert, auth, v.ValidateUser)
	path.PUT("/:ID", c.Update, auth, v.ValidateUser, v.ValidateParameterID)
	path.DELETE("/:ID", c.Delete, v.ValidateParameterID, auth)
	path.POST("/auth", c.Login, v.ValidateLogin)
	path.POST("/reset-password", c.ResetPassword(serverSocket), v.ValidateResetPassword)
	path.POST("/check-token", c.CheckToken, v.ValidateCheckToken)
	path.POST("/change-password", c.ChangePassword, v.ValidateChangePassword)
}
