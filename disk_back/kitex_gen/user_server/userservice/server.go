// Code generated by Kitex v0.10.0. DO NOT EDIT.
package userservice

import (
	server "github.com/cloudwego/kitex/server"
	user_server "github.com/cutejiuges/disk_back/kitex_gen/user_server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler user_server.UserService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler user_server.UserService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
