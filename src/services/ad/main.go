package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "demo/contracts/grpc/gen/ad/v1"

	"github.com/charmbracelet/log"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {

	port := os.Getenv("AD_SERVICE_PORT")
	if port == "" {
		port = "9555"
	}

	logger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix:          "AD",
		ReportTimestamp: true,
	})

	logger.Infof("AdService gRPC server started on port: %s", port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logger.Fatalf("TCP Listen: %v", err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	adStore := NewInMemorydStore()
	svc := NewAdGrpcService(logger, adStore)

	pb.RegisterAdServiceServer(srv, svc)
	healthpb.RegisterHealthServer(srv, svc)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	go func() {
		if err := srv.Serve(ln); err != nil {
			logger.Fatalf("Failed to serve gRPC server, err: %v", err)
		}
	}()

	<-ctx.Done()

	srv.GracefulStop()
	logger.Info("ProductCatalogService gRPC server stopped")
}
