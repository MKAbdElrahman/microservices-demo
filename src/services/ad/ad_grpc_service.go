package main

import (
	"context"
	"log"
	"os"

	"connectrpc.com/connect"

	adv1 "demo/contracts/grpc/gen/ad/v1"
)

var (
	logger = log.New(os.Stdout, "AdService: ", log.LstdFlags)
)

type AdGrpcService struct {
	adStore AdStore
}

func NewAdGrpcService(adStore AdStore) *AdGrpcService {
	return &AdGrpcService{
		adStore: adStore,
	}
}

func (s *AdGrpcService) GetAd(ctx context.Context, req *connect.Request[adv1.GetAdRequest]) (*connect.Response[adv1.GetAdResponse], error) {

	logger.Printf("received ad request (context_words=%v)", req.Msg.ContextKeys)

	allAds := s.adStore.GetAdsByCategory(req.Msg.ContextKeys)

	pbAds := make([]*adv1.Ad, len(allAds))
	for i, ad := range allAds {
		pbAds[i] = &adv1.Ad{
			RedirectUrl: ad.RedirectUrl,
			Text:        ad.Text,
		}
	}
	res := connect.NewResponse(&adv1.GetAdResponse{
		Ads: pbAds,
	})

	return res, nil
}
