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

func newStreamMetricsServer() *http.Server {

	server := &http.Server{
		Addr:    Host,
		Handler: wiredMux(),
	}

	return server
}

func wiredMux() *http.ServeMux {
	mux := http.NewServeMux()

	for endpoint, handler := range endpointMapping {
		mux.HandleFunc(endpoint, handler)
	}

	return mux
}

func RunStreamMetricsApi() *http.Server {
	server := newStreamMetricsServer()

	go func() {
		log.Printf("serving %s API at: %s\n", apiName, server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("error occurred while serving %s API: %v", apiName, err)
		}
	}()

	return server
}

func ShutDownServer(ctx context.Context, server *http.Server) error {
	if server == nil {
		return fmt.Errorf("server is nil")
	}

	log.Printf("shutting down %s API at: %s\n", apiName, server.Addr)
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("error shutting down server: %w", err)
	}

	return nil
}
