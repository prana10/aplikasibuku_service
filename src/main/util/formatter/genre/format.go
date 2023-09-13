package genre

import (
	domainGenre "service-api/src/main/app/genre/domain"
)

type GenreFormatter struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FormatGenre(genre domainGenre.Genre) GenreFormatter {
	var genreFormatter GenreFormatter

	genreFormatter.ID = genre.ID
	genreFormatter.Name = genre.Name

	return genreFormatter
}

func FormatGenres(genres []domainGenre.Genre) []GenreFormatter {
	var genresFormatter []GenreFormatter

	for _, genre := range genres {
		genreFormatter := FormatGenre(genre)
		genresFormatter = append(genresFormatter, genreFormatter)
	}

	return genresFormatter
}
