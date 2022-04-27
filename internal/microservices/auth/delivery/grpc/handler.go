package grpc

import (
	"context"
	"fmt"

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
	fmt.Println("grpc-h")
	isRegistered, err := handler.Usecase.SendCode(&models.SendCodeUcaseReq{Phone: req.GetPhone()})
	fmt.Println("grpc-eh")
	return &proto.IsRegistered{IsRegistered: isRegistered.IsRegistered}, err
}

func (handler *grpcAuthHandler) Login(ctx context.Context, req *proto.LoginReq) (*proto.UserData, error) {
	fmt.Println("grpc-h")
	userDataUcase, err := handler.Usecase.Login(&models.LoginUcaseReq{Phone: req.GetPhone(), Code: req.Code})
	fmt.Println("grpc-eh")
	if userDataUcase == nil {
		return nil, err
	}
	return &proto.UserData{Id: uint64(userDataUcase.Id), Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name, Avatar: userDataUcase.Avatar}, err
}

func (handler *grpcAuthHandler) Register(ctx context.Context, req *proto.RegisterReq) (*proto.UserData, error) {
	fmt.Println("grpc-h")
	userDataUcase, err := handler.Usecase.Register(&models.RegisterUcaseReq{Phone: req.GetPhone(), Code: req.Code, Email: req.Email, Name: req.Name})
	fmt.Println("grpc-eh")
	if userDataUcase == nil {
		return nil, err
	}
	return &proto.UserData{Id: uint64(userDataUcase.Id), Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name, Avatar: userDataUcase.Avatar}, err
}
