package repository

import (
	domainGenre "service-api/src/main/app/genre/domain"
)

type GenreRepository interface {
	GetAllGenre() ([]domainGenre.Genre, error)
	GetGenreByID(id uint) (domainGenre.Genre, error)
	InsertGenre(genre domainGenre.Genre) (domainGenre.Genre, error)
	UpdateGenre(genre domainGenre.Genre) (domainGenre.Genre, error)
	DeleteGenreByID(id uint) error
}
