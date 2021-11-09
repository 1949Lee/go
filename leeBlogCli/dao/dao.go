package dao

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"leeBlogCli/config"
	"log"
)

type DBServer struct {
	DB    *sqlx.DB
	Redis *redis.Client
}

// 打开数据库
func (s *DBServer) Open() {

	// 如果数据库连接未关闭，则关闭。
	if s.DB != nil {
		_ = s.DB.Close()
	}
	if config.ENV != "dev" {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		s.Redis = rdb
		s.InitRedis()
	}

	//连接数据库
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&multiStatements=true",
		config.DBUserName,
		config.DBUserPassword,
		config.DBIP,
		config.DBPort,
		config.DBName)
	db, err := sqlx.Connect("mysql", conStr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("数据库链接成功")
	s.DB = db
}

// 关闭数据库
func (s *DBServer) Close() {

	// 如果数据库连接未关闭，则关闭。
	if err := s.DB.Ping(); err == nil {
		log.Println("数据库关闭成功")
		_ = s.DB.Close()
	}
}
