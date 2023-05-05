package grpc

import (
	"context"
	"net"

	"github.com/christianvozar/hex-example/domain/watcher"

	"google.golang.org/grpc"
)

type server struct {
	watcher watcher.Watcher
}

func NewServer(watcher watcher.Watcher) *server {
	return &server{watcher: watcher}
}

func (s *server) TriggerUpdate(ctx context.Context) error {
	return s.watcher.ManualUpdate(ctx)
}

func StartGRPCServer(watcher watcher.Watcher, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	s := NewServer(watcher)

	grpcServer := grpc.NewServer()
	RegisterMyAppServiceServer(grpcServer, s)

	return grpcServer.Serve(lis)
}
