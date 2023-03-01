package server

import (
	"context"
	"net/http"
	"strings"

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

func customHeaderMatcher(key string) (string, bool) {
	if strings.HasPrefix(key, "X-") {
		return key, true
	}
	return runtime.DefaultHeaderMatcher(key)
}

// RunGatewayServer 启动网关服务
func RunGatewayServer(c *GatewayConfig) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(customHeaderMatcher),
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseEnumNumbers:  true,
					UseProtoNames:   true,
					EmitUnpopulated: true,
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
		panic("gateway cannot register service: " + err.Error())
	}
	addr := ":80"
	log.Infof("grpc gateway started at %s", addr)
	panic(http.ListenAndServe(addr, mux))
}
