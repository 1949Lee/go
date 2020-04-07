package dao

import (
	"leeBlogCli/definition"
	"log"
)

func (s *DBServer) GetAdminInfo() (author definition.Author) {
	rows := s.DB.QueryRowx("SELECT * FROM author where author_id=?;", 1)
	err := rows.StructScan(&author)
	if err != nil {
		log.Printf("%v", err)
	}
	return author
}
