package service_user

import (
	pb "api-gateway/internal/infra/proto/user"
	"api-gateway/internal/utils"
	"context"

	"api-gateway/internal/infra/grpc"

	"go.uber.org/zap"
)

type service struct {
	grpc pb.UserCrudClient
}

func NewServiceUser() service {
	return service{
		grpc: grpc.GetPublisherUser(),
	}
}

func (s *service) InsertUser(user *pb.User) *pb.ResponseUser {
	response, err := s.grpc.Insert(context.Background(), user)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) UpdateUser(user *pb.User) *pb.ResponseUser {
	response, err := s.grpc.Update(context.Background(), user)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) DeleteUser(ID uint64) *pb.ResponseUser {
	response, err := s.grpc.Delete(context.Background(), &pb.RequestByIdUser{Id: ID})
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}
	return response
}

func (s *service) Login(ctx context.Context, in *pb.LoginUser) (*pb.User, error) {
	return s.grpc.Login(ctx, in)
}

func (s *service) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) *pb.ResponseUser {
	response, err := s.grpc.ResetPassword(ctx, req)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}

	return response
}

func (s *service) CheckToken(ctx context.Context, req *pb.CheckTokenRequest) *pb.ResponseUser {
	response, err := s.grpc.CheckToken(ctx, req)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}

	return response
}

func (s *service) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) *pb.ResponseUser {
	response, err := s.grpc.ChangePassword(ctx, req)
	if err != nil {
		utils.LoggerZap.Error("Error en servicio", zap.String("Description", err.Error()))
	}

	return response
}
