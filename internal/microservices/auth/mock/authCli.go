package mock

// import (
// 	"context"

// 	proto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/proto"
// 	"github.com/stretchr/testify/mock"
// 	"google.golang.org/grpc"
// )

// type GrpcAuthHandler struct {
// 	mock.Mock
// }

// func (h *GrpcAuthHandler) SendCode(ctx context.Context, in *proto.SendCodeReq, opts ...grpc.CallOption) (*proto.IsRegistered, error) {
// 	return &proto.IsRegistered{IsRegistered: false}, nil
// }

// func (h *GrpcAuthHandler) Login(ctx context.Context, in *proto.LoginReq, opts ...grpc.CallOption) (*proto.UserData, error) {
// 	return &proto.UserData{
// 		Id:     1,
// 		Name:   "Name",
// 		Phone:  "79999999999",
// 		Email:  "email@mail.com",
// 		Avatar: "avatar.png"}, nil
// }

// func (h *GrpcAuthHandler) Register(ctx context.Context, in *proto.RegisterReq, opts ...grpc.CallOption) (*proto.UserData, error) {
// 	return &proto.UserData{
// 		Id:     1,
// 		Name:   "Name",
// 		Phone:  "79999999999",
// 		Email:  "email@mail.com",
// 		Avatar: "avatar.png"}, nil
// }
