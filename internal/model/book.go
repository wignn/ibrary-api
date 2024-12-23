package model

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublisedDate  string `json:"publised_date"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}


