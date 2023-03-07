package gheader

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	openidKey = "x-wx-openid"
)

// GetValue 获取header中指定key
func GetValue(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	if value, ok := md[key]; ok {
		if len(value) > 0 {
			return value[0]
		}
	}
	return ""
}

// SetWxOpenID 设置微信openid
func SetWxOpenID(ctx context.Context, openid string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, openidKey, openid)
}

// GetWxOpenID 获取微信openid
func GetWxOpenID(ctx context.Context) (string, error) {
	value := metadata.ValueFromIncomingContext(ctx, openidKey)
	if len(value) > 0 {
		return value[0], nil
	}
	return "", status.Errorf(codes.InvalidArgument, "openid empty")
}

// CheckOpenID 校验微信openid
func CheckOpenID(ctx context.Context, openid string) error {
	openid, err := GetWxOpenID(ctx)
	if err != nil {
		return err
	}
	if openid != openid {
		return status.Errorf(codes.InvalidArgument, "openid not equal")
	}
	return nil
}
