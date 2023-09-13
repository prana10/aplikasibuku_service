package impl

import (
	"errors"
	domainGenre "service-api/src/main/app/genre/domain"
	repositoryGenre "service-api/src/main/app/genre/repository"
	genreInput "service-api/src/main/util/input/genre"
)

type genreService struct {
	repository repositoryGenre.GenreRepository
}

func NewGenreService(repository repositoryGenre.GenreRepository) *genreService {
	return &genreService{repository: repository}
}

// genre all
func (s *genreService) GetAllGenre() ([]domainGenre.Genre, error) {
	genres, err := s.repository.GetAllGenre()
	if err != nil {
		return nil, err
	}

	if len(genres) == 0 {
		return nil, errors.New("genre not found")
	}

	return genres, nil
}

// genre by id
func (s *genreService) GetGenreByID(id uint) (domainGenre.Genre, error) {
	genre, err := s.repository.GetGenreByID(id)
	if err != nil {
		return domainGenre.Genre{}, err
	}

	if genre.ID == 0 {
		return genre, errors.New("genre not found")
	}

	return genre, nil
}

// insert genre
func (s *genreService) InsertGenre(genre genreInput.GenreInput) (domainGenre.Genre, error) {
	modelGenre := domainGenre.Genre{}
	modelGenre.Name = genre.Name

	newGenre, err := s.repository.InsertGenre(modelGenre)
	if err != nil {
		return domainGenre.Genre{}, err
	}

	return newGenre, nil
}

// update genre
func (s *genreService) UpdateGenre(genreId uint, genre genreInput.GenreInput) (domainGenre.Genre, error) {
	modelGenre := domainGenre.Genre{}
	modelGenre.ID = genreId
	modelGenre.Name = genre.Name

	newGenre, err := s.repository.UpdateGenre(modelGenre)
	if err != nil {
		return domainGenre.Genre{}, err
	}

	return newGenre, nil
}

// delete genre
func (s *genreService) DeleteGenreByID(id uint) error {
	err := s.repository.DeleteGenreByID(id)
	if err != nil {
		return err
	}

	return nil
}
