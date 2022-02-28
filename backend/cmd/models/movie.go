package models

import "time"

type Movie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Year        string    `json:"year"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	Rating      float32   `json:"rating"`
	MPAARating  string    `json:"mpaa_rating"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	MovieGenres []Genre   `json:"genres"`
}
