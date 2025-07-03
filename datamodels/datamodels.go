package datamodels

const (
	YouTube = StreamingPlatform("YouTube")
	Twitch  = StreamingPlatform("TwitchTV")
)

type StreamingPlatform string

type LiveStreamMetrics struct {
	Platform          StreamingPlatform `json:"Platform"`
	ChannelTitle      string            `json:"ChannelTitle"`
	StreamTitle       string            `json:"StreamTitle"`
	Thumbnail         string            `json:"Thumbnail"`
	ConcurrentViewers uint64            `json:"ConcurrentViewers"`
	Dislikes          uint64            `json:"Dislikes"`
}

type LiveStreamStatisticsQuery struct {
	StreamsStatistics []LiveStreamMetrics `json:"StreamsStatistics"`
}
