package model


type Bookmark struct {
	ID        int    `json:"id"`
	BookID    int    `json:"book_id"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}


