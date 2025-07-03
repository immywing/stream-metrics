package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"
	"stream-metrics/datamodels"
	"stream-metrics/youtubeutils"
	"text/template"
)

func ViewLiveBroadcastMetrics(w http.ResponseWriter, r *http.Request) {
	// In a real application, this would be a seperate application that called the stream-metrics API

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

	templatePath := filepath.Join("templates", "viewlivestreammetrics.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	// Render the template with the response data
	err = tmpl.Execute(w, response)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}
}
