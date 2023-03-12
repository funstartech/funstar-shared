package server

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/funstartech/funstar-shared/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

// GatewayConfig 网关服务配置
type GatewayConfig struct {
	Addr         string
	RegisterFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
}

func customHeaderMatcher(key string) (string, bool) {
	if strings.HasPrefix(key, "X-") {
		return key, true
	}
	return runtime.DefaultHeaderMatcher(key)
}

func customErrorHandler(ctx context.Context,
	mux *runtime.ServeMux, marshaler runtime.Marshaler,
	w http.ResponseWriter, r *http.Request, err error) {
	s, _ := ioutil.ReadAll(r.Body)
	log.Errorf("req: %s, err: %v", s, err)
	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}

// RunGatewayServer 启动网关服务
func RunGatewayServer(configs []*GatewayConfig) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(customErrorHandler),
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
	for _, c := range configs {
		err := c.RegisterFunc(ctx, mux, c.Addr,
			[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
		)
		if err != nil {
			panic("gateway cannot register service: " + err.Error())
		}
	}
	addr := ":80"
	log.Infof("grpc gateway started at %s", addr)
	panic(http.ListenAndServe(addr, mux))
}
