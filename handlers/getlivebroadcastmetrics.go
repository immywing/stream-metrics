package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"stream-metrics/datamodels"
	"stream-metrics/youtubeutils"
)

func GetLiveBroadcastMetrics(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing 'q' query paramater", http.StatusBadRequest)
		return
	}

	maxResults := r.URL.Query().Get("maxResults")
	if maxResults == "" {
		http.Error(w, "Missing 'maxResults' query paramater", http.StatusBadRequest)
		return
	}

	maxResultsInt, err := strconv.Atoi(maxResults)
	if err != nil {
		http.Error(w, "maxResults must be a valid integer", http.StatusBadRequest)
		return
	}

	ytService, err := youtubeutils.NewService(ctx)
	if err != nil {
		http.Error(w, "Failed to create Youtube API service", http.StatusInternalServerError)
		return
	}

	results, err := youtubeutils.QueryLiveStreamStatistics(ctx, ytService, query, int64(maxResultsInt))
	if err != nil {
		http.Error(w, "Failed to query LiveStreamStatistics", http.StatusInternalServerError)
		return
	}

	response := datamodels.LiveStreamStatisticsQuery{
		StreamsStatistics: results,
	}

	payload, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
