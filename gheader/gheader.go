package gheader

import (
	"context"

	"google.golang.org/grpc/metadata"
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
func GetWxOpenID(ctx context.Context) string {
	value := metadata.ValueFromIncomingContext(ctx, openIDKey)
	if len(value) > 0 {
		return value[0]
	}
	return ""
}
