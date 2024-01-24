package validation

import (
	objectValues "api-gateway/internal/domain/object_values"
	pb "api-gateway/internal/infra/proto/movie"
	validatorPer "api-gateway/internal/infra/validator"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateMovie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		movie := new(pb.Movie)
		_ = c.Bind(&movie)
		if err := v.Struct(movie); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := objectValues.NewResponseErrors(validatorPer.GenerateMessage(v, errs))
			return c.JSON(http.StatusBadRequest, mapErrors)
		}

		c.Set("movie", movie)
		return next(c)
	}
}
