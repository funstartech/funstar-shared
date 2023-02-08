package server

import (
	"context"
	"net/http"

	"github.com/funstartech/funstar-shared/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

// GatewayConfig 网关服务配置
type GatewayConfig struct {
	Name         string
	Addr         string
	RegisterFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
}

// RunGatewayServer 启动网关服务
func RunGatewayServer(c *GatewayConfig) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers: true,
				UseProtoNames:  true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true, // If DiscardUnknown is set, unknown fields are ignored.
			},
		},
	))
	err := c.RegisterFunc(ctx, mux, c.Addr,
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("cannot register service %s : %v", c.Name, err)
	}
	addr := ":80"
	log.Infof("grpc gateway started at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
