package dao

import (
	"database/sql"
	"leeBlogCli/definition"
	"log"
	"strconv"
	"strings"
)

// 搜索标签并附带含有对应标签的文章ID
func (s *DBServer) SelectTagsWithArticleID() (c definition.CategoryWithArticleIDListResult, err error) {
	rows, err := s.DB.Queryx("SELECT * FROM category;")
	if err != nil {
		log.Printf("%v", err)
	}
	defer rows.Close()
	c.List = make([]definition.CategoryWithTagsAndArticleID, 0)
	category := definition.CategoryWithTagsAndArticleID{}
	builder := strings.Builder{}
	// 将查询到的分类存储下来，并且存储查询该分类下标签的SQL语句
	for rows.Next() {
		err := rows.StructScan(&category)
		if err != nil {
			log.Printf("%v", err)
			c.List = append(c.List, definition.CategoryWithTagsAndArticleID{})
		}
		builder.WriteString(`SELECT t.*, GROUP_CONCAT(atr.relation_article ORDER BY atr.relation_article) AS article_ids, COUNT(atr.relation_article) as tag_length
FROM (SELECT * FROM tag where tag_category=`)
		builder.WriteString(strconv.Itoa(int(category.ID)))
		builder.WriteString(`) t 
LEFT JOIN articles_tags_relation atr ON atr.relation_tag = t.tag_id 
GROUP BY t.tag_id 
ORDER BY tag_length DESC;`)
		c.List = append(c.List, category)
	}
	tRows, err := s.DB.Queryx(builder.String())
	if err != nil {
		log.Printf("%v", err)
		return c, nil
	}
	defer tRows.Close()
	for i := range c.List {
		if c.List[i].ID == 0 {
			continue
		}
		tags := make([]definition.TagsWithArticleID, 0)
		tag := definition.TagsWithArticleID{}
		for tRows.Next() {
			var str sql.NullString
			articleIds := make([]int32, 0)
			var temTagLength int32
			err := tRows.Scan(&tag.ID, &tag.CategoryID, &tag.Name, &str, &temTagLength)
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
		c.List[i].Tags = tags
		ok := tRows.NextResultSet()
		if !ok {
			break
		}
	}
	return c, nil
}

// 根据文章ID查询文章
func (s *DBServer) GetArticleListByID(param *definition.ArticleListByIDParam) definition.ArticleListResult {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString(`SELECT * FROM (SELECT (@rownum := @rownum + 1) as i ,a.* FROM (SELECT
	a.article_id,
	a.article_ctg,
	a.article_title,
	a.article_summary,
	a.article_createtime,
	a.article_updatetime,
	c.ctg_name,
	GROUP_CONCAT(t.tag_id ORDER BY t.tag_id) AS tag_ids,
	GROUP_CONCAT(t.tag_name ORDER BY t.tag_id) AS tag_names
FROM 
`)
	sqlBuilder.WriteString(`(SELECT * FROM article WHERE article_ctg = `)
	sqlBuilder.WriteString(strconv.Itoa(int(param.CategoryID)))
	if param.ArticleIDs != "" {
		sqlBuilder.WriteString(" AND article_id IN (")
		sqlBuilder.WriteString(param.ArticleIDs)
		sqlBuilder.WriteString(")")
	}
	sqlBuilder.WriteString(` ) a 
	LEFT JOIN category c ON a.article_ctg = c.ctg_id
	LEFT JOIN articles_tags_relation r ON r.relation_article = a.article_id
	LEFT JOIN tag t ON r.relation_tag = t.tag_id
GROUP BY
	a.article_id
ORDER BY
	a.article_updatetime DESC) as a,(SELECT @rownum := -1) d) as a WHERE a.i >=? AND a.i<?;`)
	rows, err := s.DB.Queryx(sqlBuilder.String(), (param.PageIndex-1)*param.PageSize, param.PageIndex*param.PageSize)
	if err != nil {
		log.Printf("dao.GetArticleList sql报错 error：%v", err)
	}
	defer rows.Close()
	list := definition.ArticleListResult{
		List:       []definition.ArticleListResultItem{},
		IsLastPage: true,
	}
	for rows.Next() {
		tem := definition.ArticleListResultItem{
			Tags: []definition.Tag{},
		}
		rowNo := 0
		tagIDs := ""
		tagNames := ""
		err := rows.Scan(
			&rowNo,
			&tem.ID,
			&tem.CategoryID,
			&tem.Title,
			&tem.Summary,
			&tem.CreateTime,
			&tem.UpdateTime,
			&tem.CategoryName,
			&tagIDs,
			&tagNames)
		if err != nil {
			log.Printf("dao.GetArticleList 遍历SQL结果集报错 error：%v", err)
		}
		sliceTagIDs := strings.Split(tagIDs, ",")
		sliceTagNames := strings.Split(tagNames, ",")
		for i := 0; i < len(sliceTagIDs); i++ {
			temTag := definition.Tag{}
			tagID, err := strconv.Atoi(sliceTagIDs[i])
			if err != nil {
				log.Printf("dao.GetArticleList 遍历SQL结果集时转换tag_id报错 error：%v", err)
				continue
			}
			temTag.ID = int32(tagID)
			temTag.Name = sliceTagNames[i]
			temTag.CategoryID = tem.CategoryID
			tem.Tags = append(tem.Tags, temTag)

		}

		list.List = append(list.List, tem)

	}
	return list
}
