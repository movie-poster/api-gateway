package validation

import (
	"net/http"

	objectValues "api-gateway/internal/domain/object_values"
	pb "api-gateway/internal/infra/proto/user"
	validatorPer "api-gateway/internal/infra/validator"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

func ValidateLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		user := new(pb.LoginUser)

		_ = c.Bind(&user)
		if err := v.Struct(user); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := objectValues.NewResponseErrors(validatorPer.GenerateMessage(v, errs))
			return c.JSON(http.StatusBadRequest, mapErrors)
		}
		c.Set("user-data", user)
		return next(c)
	}
}

/* func ValidateVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		verify := new(objectValues.Verify)

		_ = c.Bind(&verify)
		if err := v.Struct(verify); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := map[string]any{"errors": validatorPer.GenerateMessage(v, errs)}
			return c.JSON(http.StatusBadRequest, mapErrors)
		}
		c.Set("verify", verify)
		return next(c)
	}
} */

func ValidateResetPassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		reset := new(pb.ResetPasswordRequest)

		_ = c.Bind(&reset)
		if err := v.Struct(reset); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := objectValues.NewResponseErrors(validatorPer.GenerateMessage(v, errs))
			return c.JSON(http.StatusBadRequest, mapErrors)
		}
		c.Set("reset_password", reset)
		return next(c)
	}
}

func ValidateCheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		check := new(pb.CheckTokenRequest)

		_ = c.Bind(&check)
		if err := v.Struct(check); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := objectValues.NewResponseErrors(validatorPer.GenerateMessage(v, errs))
			return c.JSON(http.StatusBadRequest, mapErrors)
		}
		c.Set("check_token", check)
		return next(c)
	}
}

func ValidateChangePassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		change := new(pb.ChangePasswordRequest)

		_ = c.Bind(&change)
		if err := v.Struct(change); err != nil {
			errs := err.(validator.ValidationErrors)
			mapErrors := objectValues.NewResponseErrors(validatorPer.GenerateMessage(v, errs))
			return c.JSON(http.StatusBadRequest, mapErrors)
		}
		c.Set("change_password", change)
		return next(c)
	}
}
