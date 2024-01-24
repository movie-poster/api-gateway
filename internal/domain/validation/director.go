package validation

import (
	objectValues "api-gateway/internal/domain/object_values"
	pb "api-gateway/internal/infra/proto/director"
	validatorPer "api-gateway/internal/infra/validator"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateDirector(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		director := new(pb.Director)
		_ = c.Bind(&director)
		if err := v.Struct(director); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := objectValues.NewResponseErrors(validatorPer.GenerateMessage(v, errs))
			return c.JSON(http.StatusBadRequest, mapErrors)
		}

		c.Set("director", director)
		return next(c)
	}
}
