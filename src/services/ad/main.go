package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"demo/contracts/grpc/gen/ad/v1/adv1connect"

	"connectrpc.com/grpchealth"
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

	adStore := NewInMemorydStore()
	adService := NewAdGrpcService(adStore)

	mux := http.NewServeMux()

	path, handler := adv1connect.NewAdServiceHandler(adService)
	checker := grpchealth.NewStaticChecker(
		adv1connect.AdServiceName,
	)
	// checker.SetStatus(adv1connect.AdServiceName, grpchealth.StatusNotServing)

	mux.Handle(grpchealth.NewHandler(checker))
	mux.Handle(path, handler)

	logger.Infof("Ad Service started, listening on %s", port)
	http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
