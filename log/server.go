package log

import (
	"context"

	"github.com/funstartech/funstar-shared/cutils"
	"google.golang.org/grpc"
)

// ServerLogInterceptor 拦截器方法
func ServerLogInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	reqJson := cutils.Obj2Json(req)
	rsp, err := handler(ctx, req)
	rspJson := cutils.Obj2Json(rsp)
	Debugf("[%v] req: %v, rsp: %v, err: %v", info.FullMethod, reqJson, rspJson, err)
	return rsp, err
}
