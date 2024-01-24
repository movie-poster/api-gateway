package validation

import (
	objectValues "api-gateway/internal/domain/object_values"
	pb "api-gateway/internal/infra/proto/actor"
	validatorPer "api-gateway/internal/infra/validator"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateActor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		actor := new(pb.Actor)
		_ = c.Bind(&actor)
		if err := v.Struct(actor); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := objectValues.NewResponseErrors(validatorPer.GenerateMessage(v, errs))
			return c.JSON(http.StatusBadRequest, mapErrors)
		}

		c.Set("actor", actor)
		return next(c)
	}
}
