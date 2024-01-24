package service_director

import (
	pb "api-gateway/internal/infra/proto/director"
	"api-gateway/internal/utils"
	"context"

	"api-gateway/internal/infra/grpc"

	"go.uber.org/zap"
)

type service struct {
	grpc pb.DirectorCrudClient
}

func NewServiceDirector() service {
	return service{
		grpc: grpc.GetPublisherDirector(),
	}
}

func (s *service) Insert(director *pb.Director) *pb.ResponseDirector {
	response, err := s.grpc.Insert(context.Background(), director)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) Update(director *pb.Director) *pb.ResponseDirector {
	response, err := s.grpc.Update(context.Background(), director)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) List(page, pageSize int) *pb.ResponseDirector {
	response, err := s.grpc.List(context.Background(), &pb.ListRequestDirector{
		Page:     uint64(page),
		PageSize: uint64(pageSize),
	})
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) Delete(ID uint64) *pb.ResponseDirector {
	response, err := s.grpc.Delete(context.Background(), &pb.RequestByIdDirector{Id: ID})
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}
