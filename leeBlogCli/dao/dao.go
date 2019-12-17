package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"leeBlogCli/config"
	"log"
)

type DBServer struct {
	DB *sqlx.DB
}

type Author struct {
	ID       string `db:"author_id"`
	Nickname string `db:"author_nickname"`
	Email    string `db:"author_email"`
	Password string `db:"author_password"`
	Motto    string `db:"author_motto"`
	IsActive string `db:"author_is_active"`
}

// 打开数据库
func (s *DBServer) Open() {

	// 如果数据库连接未关闭，则关闭。
	if s.DB != nil {
		_ = s.DB.Close()
	}

	//连接数据库
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		config.DBUserName,
		config.DBUserPassword,
		config.DBIP,
		config.DBPort,
		config.DBName)
	db, err := sqlx.Connect("mysql", conStr)
	if err != nil {
		log.Fatalln(err)
	}
	s.DB = db
}

// 关闭数据库
func (s *DBServer) Close() {

	// 如果数据库连接未关闭，则关闭。
	if err := s.DB.Ping(); err == nil {
		_ = s.DB.Close()
	}
}
