package gheader

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	openIDKey = "X-Wx-Openid"
)

// GetWxOpenID 获取微信openid
func GetWxOpenID(ctx context.Context) string {
	value := metadata.ValueFromIncomingContext(ctx, openIDKey)
	if len(value) > 0 {
		return value[0]
	}
	return ""
}
