package repository

import "github.com/wignn/library-api/internal/model"

func CreateBookmark(db *DB, bookmark *model.Bookmark) error {
	_, err := db.Exec("INSERT INTO bookmarks (book_id, user_id) VALUES ($1, $2)", bookmark.BookID, bookmark.UserID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBookmark(db *DB, id string) error {
	_, err := db.Exec("DELETE FROM bookmarks WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
