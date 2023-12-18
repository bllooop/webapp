package models

import "errors"

var NoRecord = errors.New("Запись отсутствует в базе данных")

type Movie struct {
	ID          int
	Name        string
	ReleaseDate string
	Rating      int
	Description string
}
