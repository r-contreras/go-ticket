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

	//Get genres if any
	genres, err := m.getGenresFromMovie(ctx, id)
	if err != nil {
		return nil, err
	}
	movie.MovieGenres = genres

	return &movie, nil
}

func (m *DbModel) GetAllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select
				id, title, description, year, release_date, runtime, rating, mpaa_rating
			from
				movies
			order by
				title`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var movies []*models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.Id,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}
	return movies, nil
}

func (m *DbModel) getGenresFromMovie(ctx context.Context, id int) (map[int]string, error) {
	//Scan genres related to the movie
	query := `select
				g.id, g.genre_name
			from
				genres g left join movies_genres mg on (g.id = mg.genre_id)
			where
				mg.movie_id = $1`

	rows, err := m.DB.QueryContext(ctx, query, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := make(map[int]string)

	for rows.Next() {
		var g models.Genre
		err := rows.Scan(
			&g.Id,
			&g.Name,
		)
		if err != nil {
			return nil, err
		}
		genres[g.Id] = g.Name
	}
	return genres, nil
}
