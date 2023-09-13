package genre

type GenreInput struct {
	Name string `json:"name" binding:"required"`
}
