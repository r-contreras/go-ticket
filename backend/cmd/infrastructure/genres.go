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
	//get movies related to this genre
	movies, err := m.getMoviesByGenre(ctx, id)

	if err != nil {
		return nil, err
	}

	genre.Movies = movies

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
		genres = append(genres, &genre)
	}
	return genres, nil
}

func (m *DbModel) getMoviesByGenre(ctx context.Context, id int) ([]*models.Movie, error) {

	query := `select 
				m.id, m.title, m.description, m.year, m.release_date, m.runtime, m.rating, m.mpaa_rating
			from 
				movies m left join movies_genres mg on (m.id = mg.movie_id)
			where 
				mg.genre_id = $1`

	rows, err := m.DB.QueryContext(ctx, query, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
