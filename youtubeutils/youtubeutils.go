package youtubeutils

import (
	"context"
	"fmt"
	"log"
	"stream-metrics/datamodels"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const (
	liveEventType = "live"
	videoType     = "video"
	viewCount     = "viewCount"
)

// Command line flags
var (
	YouTubeAPIKey            string
	listIDSnippet            = []string{"id", "snippet"}
	listLiveStreamStatistics = []string{"liveStreamingDetails", "snippet", "statistics"}
)

func NewService(ctx context.Context) (*youtube.Service, error) {
	return youtube.NewService(ctx, option.WithAPIKey(YouTubeAPIKey))
}

func searchLiveBroadcasts(service *youtube.Service, query string, maxResults int64) (*youtube.SearchListResponse, error) {
	return service.Search.List(listIDSnippet).Q(query).Type(videoType).EventType(liveEventType).Order(viewCount).MaxResults(maxResults).Do()
}

func getLiveStreamStatistics(service *youtube.Service, videoID string) (*youtube.VideoListResponse, error) {
	return service.Videos.List(listLiveStreamStatistics).Id(videoID).Do()
}

func getLiveStreamMetrics(service *youtube.Service, item *youtube.SearchResult) (datamodels.LiveStreamMetrics, error) {

	if item == nil || item.Id == nil || item.Snippet == nil {
		return datamodels.LiveStreamMetrics{},
			fmt.Errorf("message=%q innermessage=%q", "failed to get metrics", "insufficient data retrieved from youtube api")
	}

	if item.Id.VideoId == "" {
		return datamodels.LiveStreamMetrics{}, fmt.Errorf("message=%q innermessage=%q", "failed to get metrics", "video ID is empty")
	}

	stats, err := getLiveStreamStatistics(service, item.Id.VideoId)
	if err != nil {
		return datamodels.LiveStreamMetrics{},
			fmt.Errorf("message=%q innermessage=%w", "failed to get metrics", err)
	}

	if len(stats.Items) == 0 || stats.Items[0] == nil {
		return datamodels.LiveStreamMetrics{},
			fmt.Errorf("message=%q innermessage=%q", "failed to get metrics", "no items found")
	}

	if stats.Items[0].LiveStreamingDetails == nil {
		return datamodels.LiveStreamMetrics{},
			fmt.Errorf("message=%q innermessage=%q", "failed to get metrics", "live streaming details are nil")
	}

	return datamodels.LiveStreamMetrics{
		ChannelTitle:      item.Snippet.ChannelTitle,
		StreamTitle:       item.Snippet.Title,
		Thumbnail:         item.Snippet.Thumbnails.High.Url,
		ConcurrentViewers: stats.Items[0].LiveStreamingDetails.ConcurrentViewers,
		Dislikes:          stats.Items[0].Statistics.DislikeCount,
	}, nil
}

func QueryLiveStreamStatistics(
	ctx context.Context, service *youtube.Service, query string, maxResults int64,
) ([]datamodels.LiveStreamMetrics, error) {

	if service == nil {
		return nil, fmt.Errorf("message=%q innermessage=%q", "failed to get metrics", "youtube service is nil")
	}

	items, err := searchLiveBroadcasts(service, query, maxResults)
	if err != nil {
		return nil, err
	}

	var metrics []datamodels.LiveStreamMetrics

	for _, item := range items.Items {
		itemMetrics, err := getLiveStreamMetrics(service, item)
		if err != nil {
			log.Printf("Error retrieving metrics for item %s: %v ctx=%v", item.Id.VideoId, err, ctx)
			continue
		}
		metrics = append(metrics, itemMetrics)
	}

	return metrics, nil
}
