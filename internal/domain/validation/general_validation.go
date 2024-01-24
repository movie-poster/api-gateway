package validation

import (
	objectValues "api-gateway/internal/domain/object_values"
	"api-gateway/internal/infra/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ValidateParameterID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, err := strconv.Atoi(c.Param("ID"))
		if err != nil {
			mapErrors := objectValues.NewResponseErrors(
				[]validator.Error{
					{
						Key:     "id",
						Message: "El segmento id no contiene un número válido",
					},
				},
			)
			return c.JSON(http.StatusBadRequest, mapErrors)
		}

		c.Set("ID", ID)
		return next(c)
	}
}
