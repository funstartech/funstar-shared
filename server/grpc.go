package server

import (
	"net"

	"github.com/funstartech/funstar-shared/log"
	"google.golang.org/grpc"

	"github.com/funstartech/funstar-shared/auth"
)

// GrpcConfig grpc服务配置
type GrpcConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
}

// RunGrpcServer 启动grpc服务
func RunGrpcServer(c *GrpcConfig) {
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		panic("grpc cannot listen: " + err.Error())
	}

	var opts []grpc.ServerOption
	// 服务日志拦截器
	opts = append(opts, grpc.UnaryInterceptor(log.ServerLogInterceptor))
	// 鉴权拦截器
	if c.AuthPublicKeyFile != "" {
		in, err := auth.Interceptor(c.AuthPublicKeyFile)
		if err != nil {
			panic("cannot create auth intercept: " + err.Error())
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}

	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)

	log.Infof("[%v]server started: %v", c.Name, c.Addr)
	panic(s.Serve(lis))
}
