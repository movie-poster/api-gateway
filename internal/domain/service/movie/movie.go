package service_movie

import (
	pb "api-gateway/internal/infra/proto/movie"
	"api-gateway/internal/utils"
	"context"

	"api-gateway/internal/infra/grpc"

	"go.uber.org/zap"
)

type service struct {
	grpc pb.MovieCrudClient
}

func NewServiceMovie() service {
	return service{
		grpc: grpc.GetPublisherMovie(),
	}
}

func (s *service) Insert(movie *pb.Movie) *pb.ResponseMovie {
	response, err := s.grpc.Insert(context.Background(), movie)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) Update(movie *pb.Movie) *pb.ResponseMovie {
	response, err := s.grpc.Update(context.Background(), movie)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) List(page, pageSize int) *pb.ResponseMovie {
	response, err := s.grpc.List(context.Background(), &pb.ListRequestMovie{
		Page:     uint64(page),
		PageSize: uint64(pageSize),
	})
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) Delete(ID uint64) *pb.ResponseMovie {
	response, err := s.grpc.Delete(context.Background(), &pb.RequestByIdMovie{Id: ID})
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) GetById(ID uint64) *pb.ResponseMovie {
	response, err := s.grpc.GetById(context.Background(), &pb.RequestByIdMovie{Id: ID})
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}
