package repository

import "github.com/wignn/library-api/internal/model"

func CreateBookmark(db *DB, bookmark *model.Bookmark) error {
	smt, err := db.Prepare("INSERT INTO bookmarks (book_id, user_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer smt.Close()
	_, err = smt.Exec(bookmark.BookID, bookmark.UserID)
	return err
}

func DeleteBookmark(db *DB, id string) error {
	smt, err := db.Prepare("DELETE FROM bookmarks WHERE id = $1")

	if err != nil {
		return err
	}
	defer smt.Close()

	smt.Exec(id)
	return err
}

func GetBookmarkList(db *DB) ([]model.Bookmark, error) {
	smt, err := db.Prepare("SELECT * FROM bookmarks")
	if err != nil {
		return nil, err
	}
	defer smt.Close()

	rows, err := smt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []model.Bookmark
	for rows.Next() {
		var bookmark model.Bookmark
		if err := rows.Scan(&bookmark.ID, &bookmark.BookID, &bookmark.UserID, &bookmark.CreatedAt); err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks, err

}


func GetBookmarkById(db *DB, id string) ([]model.Bookmark, error) {
	smt, err := db.Prepare("SELECT * FROM bookmarks WHERE user_id = $1 OR book_id = $2")
	if err != nil {
		return nil, err
	}
	defer smt.Close()

	rows, err := smt.Query(id, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []model.Bookmark
	for rows.Next() {
		var bookmark model.Bookmark
		if err := rows.Scan(&bookmark.ID, &bookmark.BookID, &bookmark.UserID, &bookmark.CreatedAt); err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks, err
}