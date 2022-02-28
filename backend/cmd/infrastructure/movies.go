package infrastructure

import (
	"backend/cmd/models"
	"context"
	"time"
)

func (m *DbModel) GetMovie(id int) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, runtime, mpaa_rating, created_at, updated_at 
			from movies 
			where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie models.Movie

	err := row.Scan(
		&movie.Id,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (m *DbModel) GetAllMovies() ([]*models.Movie, error) {
	return nil, nil
}
