package service

import (
	domainGenre "service-api/src/main/app/genre/domain"
	genreInput "service-api/src/main/util/input/genre"
)

type GenreService interface {
	GetAllGenre() ([]domainGenre.Genre, error)
	GetGenreByID(id uint) (domainGenre.Genre, error)
	InsertGenre(genre genreInput.GenreInput) (domainGenre.Genre, error)
	UpdateGenre(genreId uint, genre genreInput.GenreInput) (domainGenre.Genre, error)
	DeleteGenreByID(id uint) error
}
