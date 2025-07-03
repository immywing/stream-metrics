package api

import (
	"net/http"
	"stream-metrics/handlers"
)

const (
	apiName                          = "/stream-metrics"
	getLiveBroadcastMetricsEndpoint  = "/getlivebroadcastmetrics"
	viewLiveBroadcastMetricsEndpoint = "/viewlivebroadcastmetrics"
	searchLiveBroadcastsEndpoint     = "/searchlivebroadcasts"
)

var endpointMapping = map[string]func(w http.ResponseWriter, r *http.Request){
	apiName + getLiveBroadcastMetricsEndpoint:  handlers.GetLiveBroadcastMetrics,
	apiName + viewLiveBroadcastMetricsEndpoint: handlers.ViewLiveBroadcastMetrics,
	apiName + searchLiveBroadcastsEndpoint:     handlers.SearchLiveBroadcasts,
}
