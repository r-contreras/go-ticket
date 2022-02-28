package models

import "time"

type Genre struct {
	Id        int       `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
