package main

import (
	"context"
	user_server "github.com/cutejiuges/disk_back/kitex_gen/user_server"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserSignUp implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserSignUp(ctx context.Context, req *user_server.UserSignUpRequest) (resp *user_server.UserSignUpResponse, err error) {
	// TODO: Your code here...
	return
}
