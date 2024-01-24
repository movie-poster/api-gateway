package validation

import (
	objectValues "api-gateway/internal/domain/object_values"
	pb "api-gateway/internal/infra/proto/genre"
	validatorPer "api-gateway/internal/infra/validator"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateGenre(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		genre := new(pb.Genre)
		_ = c.Bind(&genre)
		if err := v.Struct(genre); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := objectValues.NewResponseErrors(validatorPer.GenerateMessage(v, errs))
			return c.JSON(http.StatusBadRequest, mapErrors)
		}

		c.Set("genre", genre)
		return next(c)
	}
}
