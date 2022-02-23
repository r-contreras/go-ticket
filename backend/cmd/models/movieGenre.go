package models

import "time"

type MovieGenre struct {
	Id        int       `json:"id"`
	MovieId   int       `json:"movie_id"`
	GenreId   int       `json:"genre_id"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
