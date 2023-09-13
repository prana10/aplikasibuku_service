package book

type BookInput struct {
	Title       string  `json:"title" binding:"required"`
	Author      string  `json:"author" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Isbn        int     `json:"isbn" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Genre       []int   `json:"genre" binding:"required"`
}
