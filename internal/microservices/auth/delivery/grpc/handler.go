package grpc

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
	proto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (h *grpcAuthHandler) SendCode(ctx context.Context, req *proto.SendCodeReq) (*proto.IsRegistered, error) {
	fmt.Println("grpc-h")
	isRegistered, err := h.Usecase.SendCode(&models.SendCodeUcaseReq{Phone: req.GetPhone()})
	fmt.Println("grpc-eh")
	return &proto.IsRegistered{IsRegistered: isRegistered.IsRegistered}, err
}

func (h *grpcAuthHandler) Login(ctx context.Context, req *proto.LoginReq) (*proto.UserData, error) {
	fmt.Println("grpc-h")
	userDataUcase, err := h.Usecase.Login(&models.LoginUcaseReq{Phone: req.GetPhone(), Code: req.Code})
	fmt.Println("grpc-eh")
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			return &proto.UserData{}, status.Error(codes.Internal, err.Error())
		}
		return &proto.UserData{}, status.Error(codes.Code(cause.Code), err.Error())
	}
	return &proto.UserData{Id: uint64(userDataUcase.Id), Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name, Avatar: userDataUcase.Avatar}, nil
}

func (h *grpcAuthHandler) Register(ctx context.Context, req *proto.RegisterReq) (*proto.UserData, error) {
	fmt.Println("grpc-h")
	userDataUcase, err := h.Usecase.Register(&models.RegisterUcaseReq{Phone: req.GetPhone(), Code: req.Code, Email: req.Email, Name: req.Name})
	fmt.Println("grpc-eh")
	if userDataUcase == nil {
		return &proto.UserData{}, err
	}
	return &proto.UserData{Id: uint64(userDataUcase.Id), Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name, Avatar: userDataUcase.Avatar}, err
}
