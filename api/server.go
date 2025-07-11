package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

var (
	Host string
)

type StreamMetricsServer struct {
	server http.Server
}

func (s *StreamMetricsServer) Run() {

	go func() {
		log.Printf("serving %s API at: %s\n", apiName, s.server.Addr)
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("error occurred while serving %s API: %v", apiName, err)
		}
	}()

}

func (s *StreamMetricsServer) ShutDownServer(ctx context.Context) error {

	log.Printf("shutting down %s API at: %s\n", apiName, s.server.Addr)
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("error shutting down server: %w", err)
	}

	return nil
}

func NewStreamMetricsServer() *StreamMetricsServer {

	return &StreamMetricsServer{
		server: http.Server{
			Addr:    Host,
			Handler: wiredMux(),
		},
	}
}

func wiredMux() *http.ServeMux {
	mux := http.NewServeMux()

	for endpoint, handler := range endpointMapping {
		mux.HandleFunc(endpoint, handler)
	}

	return mux
}
