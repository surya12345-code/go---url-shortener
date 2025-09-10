package storage

import (
	"errors"
	"urlshortener/models"
)

var UrlDB = make(map[string]models.URL)

func GetURL(id string) (models.URL, error) {
	url, ok := UrlDB[id]
	if !ok {
		return models.URL{}, errors.New("url not found")
	}
	return url, nil
}

func SaveURL(url models.URL) {
	UrlDB[url.ID] = url
}
