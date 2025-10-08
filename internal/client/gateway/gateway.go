package gateway

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	pb "test_task/pb/storage"
)

type Gateway struct {
	client pb.StorageServiceClient
}

func NewGateway(grpcAddr string) (*Gateway, error) {
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}
	client := pb.NewStorageServiceClient(conn)
	return &Gateway{client: client}, nil
}

func (g *Gateway) Set(key, value string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := g.client.Set(ctx, &pb.SetRequest{Key: key, Value: value})
	if err != nil {
		return "", fmt.Errorf("gRPC Set error: %w", err)
	}
	return resp.Message, nil
}

func (g *Gateway) Get(key string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := g.client.Get(ctx, &pb.GetRequest{Key: key})
	if err != nil {
		return "", false, fmt.Errorf("gRPC Get error: %w", err)
	}
	return resp.Value, resp.Found, nil
}
