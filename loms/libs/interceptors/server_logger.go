package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LogServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("[gRPC Server] request %s --- %v", info.FullMethod, req)

	res, err := handler(ctx, req)
	if err != nil {
		log.Printf("[gRPC Server] error %s --- %v", info.FullMethod, err)
		return nil, err
	}

	log.Printf("[gRPC Server] responce %s --- %v", info.FullMethod, res)

	return res, nil
}
