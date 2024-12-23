package common

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

func GetTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrNoPermission(nil)
	}

	token := ""
	if val, ok := md["authorization"]; ok {
		token = val[0]
	} else {
		return "", ErrNoPermission(nil)
	}
	token = strings.TrimPrefix(token, "Bearer ")

	return token, nil
}
