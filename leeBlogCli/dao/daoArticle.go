package dao

import (
	"leeBlogCli/config"
	"leeBlogCli/definition"
	"log"
	"strconv"
	"strings"
)

// 跟句传入参数，获取文章信息
func (s *DBServer) GetArticleHeader(param *definition.GetArticleParam) {

}

func (s *DBServer) GetArticle(id int32) definition.Article {
	rows := s.DB.QueryRowx("SELECT * FROM article where article_id=?;", id)
	var article definition.Article
	_ = rows.StructScan(&article)
	return article
}

// 根据传入参数，添加文章
func (s *DBServer) InsertArticle(param *definition.SaveArticleInfo) bool {
	var summary string
	if len(param.Text) >= config.SummaryLength {
		summary = param.Text[0:config.SummaryLength]
	} else {
		summary = param.Text
	}

	tx := s.DB.MustBegin()
	defer tx.Rollback()
	_, err := tx.Exec(`INSERT INTO article (article_id,article_ctg,article_title,article_author,article_summary,article_content,article_createtime,article_updatetime) VALUES (?,?,?,?,?,?,?,?);`,
		param.ArticleHeader.ID,
		param.ArticleHeader.Category.ID,
		param.ArticleHeader.Title,
		1,       // 文章作者
		summary, // 文章摘要
		param.Content,
		param.ArticleHeader.CreateTime,
		param.ArticleHeader.CreateTime,
	)
	if err != nil {
		_ = tx.Rollback()
		log.Printf("dao.InsertArticle 插入文章报错errror:%v", err)
		return false
	}

	relationSQL := strings.Builder{}
	relationSQL.WriteString(`INSERT INTO articles_tags_relation (relation_article, relation_tag) VALUES`)
	valueSQL := make([]interface{}, 0)
	for i, tag := range param.ArticleHeader.Tags {
		if i < len(param.ArticleHeader.Tags)-1 {
			relationSQL.WriteString(`(?,?),`)
		} else {
			relationSQL.WriteString(`(?,?);`)
		}
		valueSQL = append(valueSQL, param.ArticleHeader.ID, tag.ID)
	}

	txStmt, err := tx.Prepare(relationSQL.String())
	if err != nil {
		_ = tx.Rollback()
		log.Printf("dao.InsertArticle 插入文章和标签关系SQL预编译报错errror:%v", err)
		return false
	}
	defer txStmt.Close()
	_, err = txStmt.Exec(valueSQL...)
	if err != nil {
		_ = tx.Rollback()
		log.Printf("dao.InsertArticle 插入文章和标签关系报错errror:%v", err)
		return false
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("dao.InsertArticle 报错errror:%v", err)
		return false
	} else {
		return true
	}
}

// 根据传入参数，更新对应文章
func (s *DBServer) UpdateArticle(param *definition.SaveArticleInfo) bool {
	var summary string
	if len(param.Text) >= config.SummaryLength {
		summary = param.Text[0:config.SummaryLength]
	} else {
		summary = param.Text
	}

	tx := s.DB.MustBegin()
	defer tx.Rollback()
	_, err := tx.Exec(`UPDATE article set article_ctg=?,article_title=?,article_summary=?,article_content=?,article_updatetime=? WHERE article_id=?;`,
		param.ArticleHeader.Category.ID,
		param.ArticleHeader.Title,
		summary, // 文章摘要
		param.Content,
		param.ArticleHeader.UpdateTime,
		param.ArticleHeader.ID,
	)
	if err != nil {
		_ = tx.Rollback()
		log.Printf("dao.UpdateArticle 更新文章报错errror:%v", err)
		return false
	}

	_, err = tx.Exec(`DELETE FROM articles_tags_relation WHERE relation_article=?;`, param.ArticleHeader.ID)
	if err != nil {
		_ = tx.Rollback()
		log.Printf("dao.UpdateArticle 删除原来的关系errror:%v", err)
		return false
	}

	relationSQL := strings.Builder{}
	relationSQL.WriteString(`INSERT INTO articles_tags_relation (relation_article, relation_tag) VALUES`)
	valueSQL := make([]interface{}, 0)
	for i, tag := range param.ArticleHeader.Tags {
		if i < len(param.ArticleHeader.Tags)-1 {
			relationSQL.WriteString(`(?,?),`)
		} else {
			relationSQL.WriteString(`(?,?);`)
		}
		valueSQL = append(valueSQL, param.ArticleHeader.ID, tag.ID)
	}

	txStmt, err := tx.Prepare(relationSQL.String())
	if err != nil {
		_ = tx.Rollback()
		log.Printf("dao.UpdateArticle 插入文章和标签关系SQL预编译报错errror:%v", err)
		return false
	}
	defer txStmt.Close()
	_, err = txStmt.Exec(valueSQL...)
	if err != nil {
		_ = tx.Rollback()
		log.Printf("dao.UpdateArticle 插入文章和标签关系报错errror:%v", err)
		return false
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("dao.UpdateArticle 报错errror:%v", err)
		return false
	} else {
		return true
	}
	//return true
}

// 查询文章列表
//func (s *DBServer) GetArticleList(param *definition.ArticleListParam) definition.ArticleListResult {
//	sqlBuilder := strings.Builder{}
//	sqlBuilder.WriteString(`SELECT * FROM (SELECT
//	( @rownum := @rownum + 1 ) AS i,
//	a.article_id,
//	a.article_ctg,
//	a.article_title,
//	a.article_summary,
//	a.article_createtime,
//	a.article_updatetime,
//	c.ctg_name,
//	GROUP_CONCAT(t.tag_id ORDER BY t.tag_id) AS tag_ids,
//	GROUP_CONCAT(t.tag_name ORDER BY t.tag_id) AS tag_names
//FROM
//	article a
//	LEFT JOIN category c ON a.article_ctg = c.ctg_id
//	LEFT JOIN articles_tags_relation r ON r.relation_article = a.article_id
//	LEFT JOIN tag t ON r.relation_tag = t.tag_id,
//	(SELECT @rownum := -1 as i) d `)
//	if !(param.CategoryID == 0 && param.SearchText == "" && param.TagIDs == "") {
//		condition := strings.Builder{}
//		conditions := make([]string, 0)
//		if param.CategoryID != 0 {
//			condition.WriteString("c.ctg_id = ")
//			condition.WriteString(strconv.Itoa(int(param.CategoryID)))
//			conditions = append(conditions, condition.String())
//			condition.Reset()
//		}
//		if param.SearchText != "" {
//			condition.WriteString("a.article_title LIKE '%")
//			condition.WriteString(param.SearchText)
//			condition.WriteString("%'")
//			conditions = append(conditions, condition.String())
//			condition.Reset()
//		}
//		if param.TagIDs != "" {
//			condition.WriteString("r.relation_tag IN (")
//			condition.WriteString(param.TagIDs)
//			condition.WriteString(")")
//			conditions = append(conditions, condition.String())
//			condition.Reset()
//		}
//		sqlBuilder.WriteByte('\n')
//		sqlBuilder.WriteString("WHERE ")
//		sqlBuilder.WriteString(strings.Join(conditions, " AND "))
//	}
//	sqlBuilder.WriteString(`
//GROUP BY
//	a.article_id
//ORDER BY
//	a.article_updatetime DESC) as a  WHERE a.i >=? AND a.i<?;`)
//	rows, err := s.DB.Queryx(sqlBuilder.String(), (param.PageIndex-1)*param.PageSize, param.PageIndex*param.PageSize)
//	if err != nil {
//		log.Printf("dao.GetArticleList sql报错 error：%v", err)
//	}
//	defer rows.Close()
//	list := definition.ArticleListResult{
//		List:       []definition.ArticleListResultItem{},
//		IsLastPage: true,
//	}
//	for rows.Next() {
//		tem := definition.ArticleListResultItem{
//			Tags: []definition.Tag{},
//		}
//		rowNo := 0
//		tagIDs := ""
//		tagNames := ""
//		err := rows.Scan(
//			&rowNo,
//			&tem.ID,
//			&tem.CategoryID,
//			&tem.Title,
//			&tem.Summary,
//			&tem.CreateTime,
//			&tem.UpdateTime,
//			&tem.CategoryName,
//			&tagIDs,
//			&tagNames)
//		if err != nil {
//			log.Printf("dao.GetArticleList 遍历SQL结果集报错 error：%v", err)
//		}
//		sliceTagIDs := strings.Split(tagIDs, ",")
//		sliceTagNames := strings.Split(tagNames, ",")
//		for i := 0; i < len(sliceTagIDs); i++ {
//			temTag := definition.Tag{}
//			tagID, err := strconv.Atoi(sliceTagIDs[i])
//			if err != nil {
//				log.Printf("dao.GetArticleList 遍历SQL结果集时转换tag_id报错 error：%v", err)
//				continue
//			}
//			temTag.ID = int32(tagID)
//			temTag.Name = sliceTagNames[i]
//			temTag.CategoryID = tem.CategoryID
//			tem.Tags = append(tem.Tags, temTag)
//
//		}
//
//		list.List = append(list.List, tem)
//
//	}
//	return list
//}

// 模糊查询文章列表
func (s *DBServer) GetArticleList(param *definition.ArticleListParam) definition.ArticleListResult {
	sqlBuilder := strings.Builder{}
	sqlBuilder.WriteString(`SELECT * FROM (SELECT
	( @rownum := @rownum + 1 ) AS i,
	a.article_id,
	a.article_ctg,
	a.article_title,
	a.article_summary,
	a.article_createtime,
	a.article_updatetime,
	c.ctg_name,
	GROUP_CONCAT(t.tag_id ORDER BY t.tag_id) AS tag_ids,
	GROUP_CONCAT(t.tag_name ORDER BY t.tag_id) AS tag_names
FROM`)

	// 没有查询关键字的话就是所有文章
	if param.SearchText != "" {
		//    condition.WriteString("a.article_title LIKE '%")
		//    condition.WriteString(param.SearchText)
		//    condition.WriteString("%'")
		//    conditions = append(conditions, condition.String())
		//    condition.Reset()
		sqlBuilder.WriteString(`
	(SELECT * FROM article WHERE MATCH(article.article_title,article.article_content,article.article_summary) AGAINST('`)
		sqlBuilder.WriteString(param.SearchText)
		sqlBuilder.WriteString(`')) a`)
	} else {
		sqlBuilder.WriteString(`
	article a`)
	}
	sqlBuilder.WriteString(`
	LEFT JOIN category c ON a.article_ctg = c.ctg_id
	LEFT JOIN articles_tags_relation r ON r.relation_article = a.article_id
	LEFT JOIN tag t ON r.relation_tag = t.tag_id,
	(SELECT @rownum := -1 as i) d `)
	if !(param.CategoryID == 0 && param.TagIDs == "") {
		condition := strings.Builder{}
		conditions := make([]string, 0)
		if param.CategoryID != 0 {
			condition.WriteString("c.ctg_id = ")
			condition.WriteString(strconv.Itoa(int(param.CategoryID)))
			conditions = append(conditions, condition.String())
			condition.Reset()
		}
		if param.TagIDs != "" {
			condition.WriteString("r.relation_tag IN (")
			condition.WriteString(param.TagIDs)
			condition.WriteString(")")
			conditions = append(conditions, condition.String())
			condition.Reset()
		}
		sqlBuilder.WriteByte('\n')
		sqlBuilder.WriteString("WHERE ")
		sqlBuilder.WriteString(strings.Join(conditions, " AND "))
	}
	sqlBuilder.WriteString(`
GROUP BY
	a.article_id
ORDER BY
	a.article_updatetime DESC) as a  WHERE a.i >=? AND a.i<?;`)
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

// 根据文章ID查询文章类别及所有标签
func (s *DBServer) GetArticleCategoryAndTags(id int32) (definition.Category, []definition.Tag) {
	tags := make([]definition.Tag, 0)
	category := definition.Category{}
	sqlStr := strings.Builder{}
	// 查询类别sql拼接
	sqlStr.WriteString(`SELECT
	c.*
FROM
	article a
	LEFT JOIN category c ON a.article_ctg = c.ctg_id
WHERE a.article_id=`)
	sqlStr.WriteString(strconv.Itoa(int(id)))
	sqlStr.WriteString(`;`)

	// 查询文章的所有标签的sql拼接
	sqlStr.WriteString(`SELECT
	t.*
FROM
	article a
	LEFT JOIN articles_tags_relation r ON r.relation_article = a.article_id
	LEFT JOIN tag t ON r.relation_tag = t.tag_id
WHERE a.article_id=`)
	sqlStr.WriteString(strconv.Itoa(int(id)))
	sqlStr.WriteString(`;`)
	rows, err := s.DB.Queryx(sqlStr.String())
	if err != nil {
		log.Printf("dao.GetArticleCategory 报错errror:%v", err)
		return category, tags
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&category)
		if err != nil {
			log.Printf("%v", err)
		}
	}
	ok := rows.NextResultSet()
	if ok {
		for rows.Next() {
			tag := definition.Tag{}
			err := rows.Scan(&tag.ID, &tag.CategoryID, &tag.Name)
			if err != nil {
				log.Printf("%v", err)
			}
			tags = append(tags, tag)
		}
	}
	return category, tags
}
