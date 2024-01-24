package controller_actor

import (
	sc "api-gateway/internal/domain/service/actor"
	pb "api-gateway/internal/infra/proto/actor"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Insert(c echo.Context) error {
	actor := c.Get("actor").(*pb.Actor)

	service := sc.NewServiceActor()
	response := service.Insert(actor)
	return c.JSON(int(response.GetStatus()), response)
}

func Update(c echo.Context) error {
	actor := c.Get("actor").(*pb.Actor)
	ID := c.Get("ID").(int)

	actor.Id = uint64(ID)

	service := sc.NewServiceActor()
	response := service.Update(actor)
	return c.JSON(int(response.GetStatus()), response)
}

func List(c echo.Context) error {
	page := c.Get("page").(int)
	pageSize := c.Get("pageSize").(int)

	service := sc.NewServiceActor()
	response := service.List(page, pageSize)
	return c.JSON(http.StatusOK, response)
}

func Delete(c echo.Context) error {
	ID := c.Get("ID").(int)

	service := sc.NewServiceActor()
	response := service.Delete(uint64(ID))
	return c.JSON(int(response.GetStatus()), response)
}
