package infrastructure

import (
	"backend/cmd/models"
	"context"
	"time"
)

func (m *DbModel) GetGenre(id int) (*models.Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `select id, genre_name 
				from genres
				where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var genre models.Genre

	err := row.Scan(
		&genre.Id,
		&genre.Name,
	)

	if err != nil {
		return nil, err
	}

	return &genre, nil
}

func (m *DbModel) GetAllGenres() ([]*models.Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select
				id, genre_name
			from
				genres
			order by
				genre_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*models.Genre
	for rows.Next() {
		var genre models.Genre
		err := rows.Scan(
			&genre.Id,
			&genre.Name,
		)
		if err != nil {
			return nil, err
		}
		//Maybe I'll want to implement this later
		/*Get movies if any
		genres, err := m.getgenreGenres(ctx, genre.Id)
		if err != nil {
			return nil, err
		}
		genre.genreGenres = genres
		*/
		genres = append(genres, &genre)
	}
	return genres, nil
}
