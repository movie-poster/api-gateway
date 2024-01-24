package validation

import (
	objectValues "api-gateway/internal/domain/object_values"
	pb "api-gateway/internal/infra/proto/user"
	validatorPer "api-gateway/internal/infra/validator"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		user := new(pb.User)
		_ = c.Bind(&user)
		if err := v.Struct(user); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := objectValues.NewResponseErrors(validatorPer.GenerateMessage(v, errs))
			return c.JSON(http.StatusBadRequest, mapErrors)
		}

		c.Set("user_entity", user)
		return next(c)
	}
}
