package main

import (
	"flag"
	"fmt"
)

// Command line flags
var (
	youtubeAPIKey string
)

func main() {
	flag.StringVar(&youtubeAPIKey, "youtube-api-key", "", "API key for YouTube API")
	flag.Parse()
	fmt.Println(youtubeAPIKey)
}
