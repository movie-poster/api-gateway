package controller_genre

import (
	sc "api-gateway/internal/domain/service/genre"
	pb "api-gateway/internal/infra/proto/genre"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Insert(c echo.Context) error {
	genre := c.Get("genre").(*pb.Genre)

	service := sc.NewServiceGenre()
	response := service.Insert(genre)
	return c.JSON(int(response.GetStatus()), response)
}

func List(c echo.Context) error {
	page := c.Get("page").(int)
	pageSize := c.Get("pageSize").(int)

	service := sc.NewServiceGenre()
	response := service.List(page, pageSize)
	return c.JSON(http.StatusOK, response)
}

func Delete(c echo.Context) error {
	ID := c.Get("ID").(int)

	service := sc.NewServiceGenre()
	response := service.Delete(uint64(ID))
	return c.JSON(int(response.GetStatus()), response)
}
