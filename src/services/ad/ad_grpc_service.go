package main

import (
	"context"

	pb "demo/contracts/grpc/gen/ad/v1"

	"github.com/charmbracelet/log"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type AdGrpcService struct {
	adStore AdStore
	pb.UnimplementedAdServiceServer
	logger *log.Logger
}

func NewAdGrpcService(logger *log.Logger, adStore AdStore) *AdGrpcService {
	return &AdGrpcService{
		adStore: adStore,
		logger:  logger,
	}
}

func (s *AdGrpcService) GetAds(ctx context.Context, req *pb.GetAdRequest) (*pb.GetAdResponse, error) {

	s.logger.Infof("received ad request (context_words=%v)", req.ContextKeys)

	allAds := s.adStore.GetAdsByCategory(req.ContextKeys)

	pbAds := make([]*pb.Ad, len(allAds))
	for i, ad := range allAds {
		pbAds[i] = &pb.Ad{
			RedirectUrl: ad.RedirectUrl,
			Text:        ad.Text,
		}
	}
	res := &pb.GetAdResponse{
		Ads: pbAds,
	}

	return res, nil
}

func (p *AdGrpcService) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (p *AdGrpcService) Watch(req *healthpb.HealthCheckRequest, ws healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
