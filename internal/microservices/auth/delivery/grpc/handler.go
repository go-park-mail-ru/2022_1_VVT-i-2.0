package grpc

import (
	"context"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
	proto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcAuthHandler struct {
	Ucase auth.Ucase
	proto.UnimplementedAuthServiceServer
}

func NewAuthHandler(userUcase auth.Ucase) *grpcAuthHandler {
	return &grpcAuthHandler{
		Ucase: userUcase,
	}
}

func (h *grpcAuthHandler) SendCode(ctx context.Context, req *proto.SendCodeReq) (*proto.IsRegistered, error) {
	isRegistered, err := h.Ucase.SendCode(&models.SendCodeUcaseReq{Phone: req.GetPhone()})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			return &proto.IsRegistered{}, status.Error(codes.Internal, err.Error())
		}
		return &proto.IsRegistered{}, status.Error(codes.Code(cause.Code), err.Error())
	}
	return &proto.IsRegistered{IsRegistered: isRegistered.IsRegistered}, nil
}

func (h *grpcAuthHandler) Login(ctx context.Context, req *proto.LoginReq) (*proto.UserData, error) {
	userDataUcase, err := h.Ucase.Login(&models.LoginUcaseReq{Phone: req.GetPhone(), Code: req.Code})
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
	userDataUcase, err := h.Ucase.Register(&models.RegisterUcaseReq{Phone: req.GetPhone(), Code: req.Code, Email: req.Email, Name: req.Name})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			return &proto.UserData{}, status.Error(codes.Internal, err.Error())
		}
		return &proto.UserData{}, status.Error(codes.Code(cause.Code), err.Error())
	}
	return &proto.UserData{Id: uint64(userDataUcase.Id), Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name, Avatar: userDataUcase.Avatar}, nil
}
