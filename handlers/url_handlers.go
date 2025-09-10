package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"urlshortener/services"
	"urlshortener/storage"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello surya")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Url string `json:"url"`
	}
	_ = json.NewDecoder(r.Body).Decode(&data)

	shortID := services.CreateURL(data.Url)
	resp := map[string]string{
		"original_url": data.Url,
		"short_url":    "http://localhost:8080/bitly/" + shortID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/bitly/"):]
	url, err := storage.GetURL(id)
	if err != nil {
		http.Error(w, "invalid request", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
