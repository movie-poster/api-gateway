package controller_director

import (
	sc "api-gateway/internal/domain/service/director"
	pb "api-gateway/internal/infra/proto/director"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Insert(c echo.Context) error {
	director := c.Get("director").(*pb.Director)

	service := sc.NewServiceDirector()
	response := service.Insert(director)
	return c.JSON(int(response.GetStatus()), response)
}

func Update(c echo.Context) error {
	director := c.Get("director").(*pb.Director)
	ID := c.Get("ID").(int)

	director.Id = uint64(ID)

	service := sc.NewServiceDirector()
	response := service.Update(director)
	return c.JSON(int(response.GetStatus()), response)
}

func List(c echo.Context) error {
	page := c.Get("page").(int)
	pageSize := c.Get("pageSize").(int)

	service := sc.NewServiceDirector()
	response := service.List(page, pageSize)
	return c.JSON(http.StatusOK, response)
}

func Delete(c echo.Context) error {
	ID := c.Get("ID").(int)

	service := sc.NewServiceDirector()
	response := service.Delete(uint64(ID))
	return c.JSON(int(response.GetStatus()), response)
}
