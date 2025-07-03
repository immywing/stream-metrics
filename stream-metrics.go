package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"stream-metrics/api"
	"stream-metrics/youtubeutils"
)

func parseFlags() {
	flag.StringVar(&youtubeutils.YouTubeAPIKey, "youtube-api-key", "", "API key for YouTube API")
	flag.StringVar(&api.Host, "host", "", "The host for the stream metrics API")
	flag.Parse()
}

func validateCLIArguments() error {
	if youtubeutils.YouTubeAPIKey == "" {
		return fmt.Errorf("YouTube API key is required. Use -youtube-api-key flag")
	}
	if api.Host == "" {
		return fmt.Errorf("Host is required. Use -host flag")
	}
	return nil
}

func main() {

	parseFlags()
	fatalError := validateCLIArguments()

	if fatalError != nil {
		fmt.Printf("Error: %v\n", fatalError)
		os.Exit(1)
	}

	server := api.RunStreamMetricsApi()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
	_ = api.ShutDownServer(context.Background(), server)

}
