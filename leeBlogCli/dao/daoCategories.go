package dao

import (
	"database/sql"
	"leeBlogCli/definition"
	"log"
	"strconv"
	"strings"
)

// 搜索标签并附带含有对应标签的文章ID
func (s *DBServer) SelectTagsWithArticleID() (c []definition.CategoryWithTagsAndArticleID, err error) {
	rows, err := s.DB.Queryx("SELECT * FROM category;")
	if err != nil {
		log.Printf("%v", err)
	}
	defer rows.Close()
	c = make([]definition.CategoryWithTagsAndArticleID, 0)
	category := definition.CategoryWithTagsAndArticleID{}
	builder := strings.Builder{}
	// 将查询到的分类存储下来，并且存储查询该分类下标签的SQL语句
	for rows.Next() {
		err := rows.StructScan(&category)
		if err != nil {
			log.Printf("%v", err)
			c = append(c, definition.CategoryWithTagsAndArticleID{})
		}
		builder.WriteString(`SELECT t.*, GROUP_CONCAT(atr.relation_article ORDER BY atr.relation_article) AS article_ids
FROM (SELECT * FROM tag where tag_category=`)
		builder.WriteString(strconv.Itoa(int(category.ID)))
		builder.WriteString(`) t 
LEFT JOIN articles_tags_relation atr ON atr.relation_tag = t.tag_id 
GROUP BY t.tag_id;`)
		c = append(c, category)
	}
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
		tags := make([]definition.TagsWithArticleID, 0)
		tag := definition.TagsWithArticleID{}
		for tRows.Next() {
			var str sql.NullString
			articleIds := make([]int32, 0)
			err := tRows.Scan(&tag.ID, &tag.CategoryID, &tag.Name, &str)
			if err != nil {
				log.Printf("%v", err)
			}
			if str.Valid {
				sliceArticleIds := strings.Split(str.String, ",")
				for _, id := range sliceArticleIds {
					articleID, err := strconv.Atoi(id)
					if err != nil {
						log.Printf("dao.SelectTagsWithArticleID 遍历SQL结果集时转换article_ids报错 error：%v", err)
						continue
					}
					articleIds = append(articleIds, int32(articleID))
				}
			} else {
				articleIds = nil
			}
			tag.ArticleIDs = articleIds

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
