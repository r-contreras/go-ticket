package models

import "time"

type MovieGenre struct {
	Id        int       `json:"-"`
	MovieId   int       `json:"-"`
	GenreId   int       `json:"genre_id"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
