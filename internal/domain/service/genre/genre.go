package service_genre

import (
	pb "api-gateway/internal/infra/proto/genre"
	"api-gateway/internal/utils"
	"context"

	"api-gateway/internal/infra/grpc"

	"go.uber.org/zap"
)

type service struct {
	grpc pb.GenreCrudClient
}

func NewServiceGenre() service {
	return service{
		grpc: grpc.GetPublisherGenre(),
	}
}

func (s *service) Insert(genre *pb.Genre) *pb.ResponseGenre {
	response, err := s.grpc.Insert(context.Background(), genre)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) List(page, pageSize int) *pb.ResponseGenre {
	response, err := s.grpc.List(context.Background(), &pb.ListRequestGenre{
		Page:     uint64(page),
		PageSize: uint64(pageSize),
	})
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) Delete(ID uint64) *pb.ResponseGenre {
	response, err := s.grpc.Delete(context.Background(), &pb.RequestByIdGenre{Id: ID})
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}
