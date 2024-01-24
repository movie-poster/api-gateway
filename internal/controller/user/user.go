package controller_user

import (
	objectValues "api-gateway/internal/domain/object_values"
	su "api-gateway/internal/domain/service/user"
	"api-gateway/internal/infra/jwt"
	pb "api-gateway/internal/infra/proto/user"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
)

func Insert(c echo.Context) error {
	user := c.Get("user_entity").(*pb.User)

	service := su.NewServiceUser()
	response := service.InsertUser(user)

	return c.JSON(int(response.GetStatus()), response)
}

func Update(c echo.Context) error {
	user := c.Get("user_entity").(*pb.User)
	ID := c.Get("ID").(int)

	user.Id = uint64(ID)

	service := su.NewServiceUser()
	response := service.UpdateUser(user)

	return c.JSON(int(response.GetStatus()), response)
}

func Delete(c echo.Context) error {
	ID := c.Get("ID").(int)

	service := su.NewServiceUser()
	response := service.DeleteUser(uint64(ID))
	return c.JSON(int(response.GetStatus()), response)
}

func Login(c echo.Context) error {
	user := c.Get("user-data").(*pb.LoginUser)
	service := su.NewServiceUser()
	res, err := service.Login(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"auth":    false,
			"title":   "DenyAuth",
			"message": "usuario no autenticado, no encontrado",
		})
	}
	jsW := jwt.NewJwtClient()
	token, err := jsW.GenerateToken(objectValues.JwtEntity{
		Name:     res.GetName(),
		Surname:  res.GetSurname(),
		Email:    res.GetEmail(),
		UserId:   res.GetId(),
		NickName: res.GetNickName(),
	})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"auth":    false,
			"title":   "DenyAuth",
			"message": "error 403, internal server error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"auth":    true,
		"title":   "Auth",
		"message": "Usuario logeado satisfactoriamente",
		"token":   token,
	})
}

func ResetPassword(server *socketio.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		reset := c.Get("reset_password").(*pb.ResetPasswordRequest)

		jsW := jwt.NewJwtClient()
		token, err := jsW.GenerateTokenResetPassword(reset.GetNickname())
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"title":   "Proceso no existoso",
				"message": "Inténtalo más tarde por favor",
				"status":  http.StatusBadRequest,
			})
		}
		reset.Token = token

		service := su.NewServiceUser()
		response := service.ResetPassword(c.Request().Context(), reset)

		return c.JSON(int(response.GetStatus()), response)
	}
}

func CheckToken(c echo.Context) error {
	check := c.Get("check_token").(*pb.CheckTokenRequest)

	instance := jwt.NewJwtClient()
	valid := instance.VerifyTokenResetPassword(check.GetToken())
	if !valid {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"title":   "Proceso no existoso",
			"message": "El token es inválido o ha expirado. Te pedimos que solicites nuevamente un cambio de contraseña. Recuerda que tienes dos horas para poder realizar el proceso",
			"status":  http.StatusBadRequest,
			"is_ok":   false,
		})
	}

	service := su.NewServiceUser()
	response := service.CheckToken(c.Request().Context(), check)
	return c.JSON(int(response.GetStatus()), response)
}

func ChangePassword(c echo.Context) error {
	change := c.Get("change_password").(*pb.ChangePasswordRequest)

	instance := jwt.NewJwtClient()
	valid := instance.VerifyTokenResetPassword(change.GetToken())
	if !valid {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"title":   "Proceso no existoso",
			"message": "El token es inválido o ha expirado. Te pedimos que solicites nuevamente un cambio de contraseña. Recuerda que tienes dos horas para poder realizar el proceso",
			"status":  http.StatusBadRequest,
			"is_ok":   false,
		})
	}

	service := su.NewServiceUser()
	response := service.ChangePassword(c.Request().Context(), change)
	return c.JSON(int(response.GetStatus()), response)
}
