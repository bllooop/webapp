package psql

import (
	"database/sql"
	"errors"

	"go.com/movies/cmd/web/pkg/models"
)

type MovieModel struct {
	DB *sql.DB
}

func (m *MovieModel) Insert(name, releaseDate, rating, description string) (int, error) {
	var id int
	statement := `INSERT INTO movies (name, releaseDate, rating, description) VALUES (?,?,?,?) RETURNING id`
	err := m.DB.QueryRow(statement, name, releaseDate, rating, description).Scan(&id)
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (m *MovieModel) Get(id int) (*models.Movie, error) {
	statement := `SELECT FROM name, releaseDate, rating, description FROM movies WHERE id = ?`
	res := m.DB.QueryRow(statement, id)
	k := &models.Movie{}
	err := res.Scan(&k.Name, &k.ReleaseDate, &k.Rating, &k.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.NoRecord
		} else {
			return nil, err
		}
	}
	return k, nil
}
func (m *MovieModel) LastTwenty() ([]*models.Movie, error) {
	statement := `SELECT FROM id, name, releaseDate, rating, description FROM movies ORDER BY id DESC LIMIT 20`
	result, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var movies []*models.Movie
	for result.Next() {
		k := &models.Movie{}
		err := result.Scan(&k.ID, &k.Name, &k.ReleaseDate, &k.Rating, &k.Description)
		if err != nil {
			return nil, err
		}
		movies = append(movies, k)
	}
	if err = result.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}
