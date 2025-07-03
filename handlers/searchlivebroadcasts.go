package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func SearchLiveBroadcasts(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("templates", "searchlivebroadcasts.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
