package services

import (
	"time"
	"urlshortener/models"
	"urlshortener/storage"
	"urlshortener/utils"
)

func CreateURL(originalURL string) string {
	id := utils.GenerateShortID(originalURL)

	for {
		if existing, ok := storage.UrlDB[id]; ok {
			if existing.OriginalURL == originalURL {
				return id // Reuse same short ID
			}
			// Collision → regenerate with timestamp
			id = utils.GenerateShortID(originalURL + time.Now().String())
		} else {
			break
		}
	}

	newURL := models.URL{
		ID:           id,
		OriginalURL:  originalURL,
		Shorturl:     id,
		CreationDate: time.Now(),
	}

	storage.SaveURL(newURL)
	return id
}
