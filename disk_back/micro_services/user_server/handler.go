package main

import (
	"context"
	user_server "github.com/cutejiuges/disk_back/kitex_gen/user_server"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/handler"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserSignUp implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserSignUp(ctx context.Context, req *user_server.UserSignUpRequest) (resp *user_server.UserSignUpResponse, err error) {
	return handler.NewUserSignUpHandler(ctx, req).Handle()
}

// GetEmailVerifyCode implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetEmailVerifyCode(ctx context.Context, req *user_server.GetEmailVerifyCodeRequest) (resp *user_server.GetEmailVerifyCodeResponse, err error) {
	return handler.NewGetEmailVerifyCodeHandler(ctx, req).Handle()
}

// UserSignIn implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserSignIn(ctx context.Context, req *user_server.UserSignInRequest) (resp *user_server.UserSignInResponse, err error) {
	// TODO: Your code here...
	return
}
