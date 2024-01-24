package custom_middleware

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func ValidatePageQueryParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParams().Get("page"))
		if err != nil {
			page = 0
		}

		c.Set("page", page)
		return next(c)
	}
}

func ValidatePageSizeQueryParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		pageSize, err := strconv.Atoi(c.QueryParams().Get("pageSize"))
		if err != nil {
			pageSize = 10
		}

		c.Set("pageSize", pageSize)
		return next(c)
	}
}
