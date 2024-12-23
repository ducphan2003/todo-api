package common

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func GetMetadataFromContext(ctx context.Context) (map[string]string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, false
	}

	metadataMap := make(map[string]string)
	for key, values := range md {
		if len(values) > 0 {
			metadataMap[key] = values[0]
		}
	}

	return metadataMap, true
}
