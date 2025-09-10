package models

import "time"

type URL struct {
	ID           string
	OriginalURL  string
	Shorturl     string
	CreationDate time.Time
}
