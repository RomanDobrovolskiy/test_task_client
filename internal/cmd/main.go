package main

import (
	"fmt"
	"log"
	"net/http"
	gw "test_task/internal/client/gateway"
	handler "test_task/internal/client/http"
)

func main() {
	// В идеальном варианте вынести в конфиг, сократил для экономии времени
	grpcAddr := "localhost:50051" // адрес gRPC-сервера
	httpAddr := ":8080"           // адрес шлюза

	gateway, err := gw.NewGateway(grpcAddr)
	if err != nil {
		log.Fatalf("failed to init gateway: %v", err)
	}

	httpHandler := handler.NewHTTPHandler(gateway)

	fmt.Printf("✅ HTTP Gateway started on %s → gRPC %s\n", httpAddr, grpcAddr)
	if err := http.ListenAndServe(httpAddr, httpHandler); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
