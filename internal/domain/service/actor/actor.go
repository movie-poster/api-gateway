package service_actor

import (
	pb "api-gateway/internal/infra/proto/actor"
	"api-gateway/internal/utils"
	"context"

	"api-gateway/internal/infra/grpc"

	"go.uber.org/zap"
)

type service struct {
	grpc pb.ActorCrudClient
}

func NewServiceActor() service {
	return service{
		grpc: grpc.GetPublisherActor(),
	}
}

func (s *service) Insert(actor *pb.Actor) *pb.ResponseActor {
	response, err := s.grpc.Insert(context.Background(), actor)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) Update(actor *pb.Actor) *pb.ResponseActor {
	response, err := s.grpc.Update(context.Background(), actor)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) List(page, pageSize int) *pb.ResponseActor {
	response, err := s.grpc.List(context.Background(), &pb.ListRequestActor{
		Page:     uint64(page),
		PageSize: uint64(pageSize),
	})
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) Delete(ID uint64) *pb.ResponseActor {
	response, err := s.grpc.Delete(context.Background(), &pb.RequestByIdActor{Id: ID})
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}
