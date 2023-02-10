package log

import (
	"context"

	"github.com/funstartech/funstar-shared/cutils"
	"google.golang.org/grpc"
)

// ServerLogInterceptor 拦截器方法
func ServerLogInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	rsp, err := handler(ctx, req)
	Debugf("req: %v, rsp: %v, err: %v", cutils.Obj2Json(req), cutils.Obj2Json(rsp), err)
	return rsp, err
}
