package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func tokenFromCtx(c context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "unauthenticated")
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", unauthenticated
	}
	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn == "" {
		return "", unauthenticated
	}
	return tkn, nil
}

func ctxWithAccountID(c context.Context, aid string) context.Context {
	return context.WithValue(c, accountIDKey, aid)
}

// AccountIDFromCtx 从ctx中取出账号id
func AccountIDFromCtx(c context.Context) (string, error) {
	v := c.Value(accountIDKey)
	aid, ok := v.(string)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}
