package model


type Genre struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type BookGenre struct {
	BookID  int `json:"book_id"`
	GenreID int `json:"genre_id"`
}

