package dao

import (
	"leeBlogCli/config"
	"leeBlogCli/definition"
	"log"
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
	//if
	return true
}
