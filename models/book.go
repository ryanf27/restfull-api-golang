package models

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Year        int    `json:"year"`
}

var Books = []Book{
	{
		ID:          1,
		Title:       "Book 1",
		Author:      "Author 1",
		Description: "Description 1",
		Year:        2020,
	},
	{
		ID:          2,
		Title:       "Book 2",
		Author:      "Author 2",
		Description: "Description 2",
		Year:        2021,
	},
}