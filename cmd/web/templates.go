package main

import "go.com/movies/cmd/web/pkg/models"

type movieData struct {
	Movie  *models.Movie
	Movies []*models.Movie
}
