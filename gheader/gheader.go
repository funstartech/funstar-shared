package gheader

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	openIDKey = "x-wx-openid"
	sourceKey = "x-wx-source"
)

// GetWxOpenID 获取微信openid
func GetWxOpenID(ctx context.Context) string {
	value := metadata.ValueFromIncomingContext(ctx, openIDKey)
	if len(value) > 0 {
		return value[0]
	}
	return ""
}

// GetWxSource 获取微信source
func GetWxSource(ctx context.Context) string {
	value := metadata.ValueFromIncomingContext(ctx, sourceKey)
	if len(value) > 0 {
		return value[0]
	}
	return ""
}
