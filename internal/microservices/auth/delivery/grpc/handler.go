package grpc

import (
	"context"

	proto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/proto/auth"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
)

type grpcAuthHandler struct {
	Usecase auth.Usecase
	proto.UnimplementedAuthServiceServer
}

func NewAuthHandler(userUsecase auth.Usecase) *grpcAuthHandler {
	return &grpcAuthHandler{
		Usecase: userUsecase,
	}
}

func (handler *grpcAuthHandler) SendCode(ctx context.Context, req *proto.SendCodeReq) (*proto.IsRegistered, error) {
	isRegistered, err := handler.Usecase.SendCode(&models.SendCodeUcaseReq{Phone: req.GetPhone()})
	return &proto.IsRegistered{IsRegistered: isRegistered.IsRegistered}, err
}

// func (handler *grpcAuthHandler) SignUp(ctx context.Context, userInput *proto.UserForReg) (*proto.UserId, error) {
// 	userID, err := handler.userUseCase.SignUp(models.GrpcUserDataForRegToModel(userInput))
// 	if err != nil {
// 		return models.ModelUserIdToGrpc(models.UserID{UserId: userID}), err
// 	}
// 	return models.ModelUserIdToGrpc(models.UserID{UserId: userID}), err
// }
