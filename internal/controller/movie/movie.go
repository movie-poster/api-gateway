package controller_movie

import (
	sc "api-gateway/internal/domain/service/movie"
	pb "api-gateway/internal/infra/proto/movie"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Insert(c echo.Context) error {
	movie := c.Get("movie").(*pb.Movie)

	service := sc.NewServiceMovie()
	response := service.Insert(movie)
	return c.JSON(int(response.GetStatus()), response)
}

func Update(c echo.Context) error {
	movie := c.Get("movie").(*pb.Movie)
	ID := c.Get("ID").(int)

	movie.Id = uint64(ID)

	service := sc.NewServiceMovie()
	response := service.Update(movie)
	return c.JSON(int(response.GetStatus()), response)
}

func List(c echo.Context) error {
	page := c.Get("page").(int)
	pageSize := c.Get("pageSize").(int)

	service := sc.NewServiceMovie()
	response := service.List(page, pageSize)
	return c.JSON(http.StatusOK, response)
}

func Delete(c echo.Context) error {
	ID := c.Get("ID").(int)

	service := sc.NewServiceMovie()
	response := service.Delete(uint64(ID))
	return c.JSON(int(response.GetStatus()), response)
}

func GetById(c echo.Context) error {
	ID := c.Get("ID").(int)

	service := sc.NewServiceMovie()
	response := service.GetById(uint64(ID))

	return c.JSON(int(response.GetStatus()), response)
}
