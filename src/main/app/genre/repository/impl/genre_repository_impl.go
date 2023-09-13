package impl

import (
	domainGenre "service-api/src/main/app/genre/domain"

	"gorm.io/gorm"
)

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) *genreRepository {
	return &genreRepository{db: db}
}

// genre all
func (r *genreRepository) GetAllGenre() ([]domainGenre.Genre, error) {
	var genres []domainGenre.Genre
	result := r.db.Find(&genres)
	if result.Error != nil {
		return nil, result.Error
	}
	return genres, nil
}

// genre by id
func (r *genreRepository) GetGenreByID(id uint) (domainGenre.Genre, error) {
	var genre domainGenre.Genre
	result := r.db.First(&genre, id)
	if result.Error != nil {
		return domainGenre.Genre{}, result.Error
	}
	return genre, nil
}

// insert genre
func (r *genreRepository) InsertGenre(genre domainGenre.Genre) (domainGenre.Genre, error) {
	result := r.db.Create(&genre)
	if result.Error != nil {
		return domainGenre.Genre{}, result.Error
	}
	return genre, nil
}

// update genre
func (r *genreRepository) UpdateGenre(genre domainGenre.Genre) (domainGenre.Genre, error) {
	result := r.db.Save(&genre)
	if result.Error != nil {
		return domainGenre.Genre{}, result.Error
	}
	return genre, nil
}

// delete genre
func (r *genreRepository) DeleteGenreByID(id uint) error {
	var genre domainGenre.Genre
	result := r.db.Delete(&genre, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
