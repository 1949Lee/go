package dao

import (
	"database/sql"
	"log"
)

func (s *DBServer) NextArticleID() (ID int32) {
	rows := s.DB.QueryRow("SELECT MAX(article_id) + 1 as ID FROM article;")
	var col sql.NullInt32
	err := rows.Scan(&col)
	if err != nil {
		log.Printf("%v", err)
	}
	if col.Valid {
		ID = col.Int32
	} else {
		ID = 1
	}
	return ID
}
