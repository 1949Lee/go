package dao

import (
	"database/sql"
	"leeBlogCli/definition"
	"log"
	"strconv"
	"strings"
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

func (s *DBServer) SelectTagsGroupByCategory() (c []definition.CategoryWithTags, err error) {
	rows, err := s.DB.Queryx("SELECT * FROM category;")
	if err != nil {
		log.Printf("%v", err)
	}
	c = make([]definition.CategoryWithTags, 0)
	category := definition.CategoryWithTags{}
	builder := strings.Builder{}
	// 将查询到的分类存储下来，并且存储查询该分类下标签的SQL语句
	for rows.Next() {
		err := rows.StructScan(&category)
		if err != nil {
			log.Printf("%v", err)
			c = append(c, definition.CategoryWithTags{})
		}
		builder.WriteString("SELECT * FROM tag where tag_category=")
		builder.WriteString(strconv.Itoa(int(category.ID)))
		builder.WriteString("; ")
		c = append(c, category)
	}
	_ = rows.Close()
	tRows, err := s.DB.Queryx(builder.String())
	if err != nil {
		log.Printf("%v", err)
		return c, nil
	}
	defer tRows.Close()
	for i := range c {
		if c[i].ID == 0 {
			continue
		}
		tags := make([]definition.Tag, 0)
		tag := definition.Tag{}
		for tRows.Next() {
			err := tRows.StructScan(&tag)
			if err != nil {
				log.Printf("%v", err)
			}
			tags = append(tags, tag)
		}
		if len(tags) == 0 {
			tags = nil
		}
		c[i].Tags = tags
		ok := tRows.NextResultSet()
		if !ok {
			break
		}
	}
	return c, nil
}