package dao

import (
	"database/sql"
	"leeBlogCli/definition"
	"log"
)

// 获取管理员信息
func (s *DBServer) GetAdminInfo() (author definition.Author) {
	rows := s.DB.QueryRowx("SELECT * FROM author where author_id=?;", 1)
	err := rows.StructScan(&author)
	if err != nil {
		log.Printf("%v", err)
	}
	return author
}

// 根据email查询用户信息
func (s *DBServer) GetAuthorByEmail(email string) (author definition.Author) {
	rows := s.DB.QueryRowx("SELECT * FROM author where author_email=?;", email)
	err := rows.StructScan(&author)

	if err == sql.ErrNoRows {
		author.ID = -1
	} else if err != nil {
		log.Printf("%v", err)
		author.ID = -2
	}
	return author
}

// 根据email和leeToken设置登录。
func (s *DBServer) UpdateAuthorToken(email string, token string) error {
	_, err := s.DB.Exec("UPDATE author set author_lee_token=? where author_email=?", token, email)
	if err != nil {
		log.Printf("%v", err)
	}
	return err
}
