package pkg

import (
	"os"

	"google.golang.org/grpc"
)

func NewAuthClientConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(os.Getenv("AUTH_GRPC_ADDR"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
