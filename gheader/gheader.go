package gheader

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	openIDKey = "x-wx-openid"
)

// GetValue 获取header中指定key
func GetValue(ctx context.Context, key string) string {
	md, ok := metadata.FromOutgoingContext(ctx)
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

// GetWxOpenID 获取微信openid
func GetWxOpenID(ctx context.Context) (string, error) {
	value := metadata.ValueFromIncomingContext(ctx, openIDKey)
	if len(value) > 0 {
		return value[0], nil
	}
	return "", status.Errorf(codes.InvalidArgument, "openid empty")
}

// CheckOpenID 校验微信openid
func CheckOpenID(ctx context.Context, openID string) error {
	openid, err := GetWxOpenID(ctx)
	if err != nil {
		return err
	}
	if openid != openID {
		return status.Errorf(codes.InvalidArgument, "openid not equal")
	}
	return nil
}
