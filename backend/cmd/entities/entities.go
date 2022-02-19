package entities

type Movie struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Year string `json:"year"`
}

type Category struct {
	Id int `json:"id"`
	Name string `json:"name"`
}