package dao

import (
	"leeBlogCli/definition"
	"log"
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
	_, err := s.DB.Exec(`INSERT INTO article (article_id,article_ctg,article_title,article_author,article_summary,article_content,article_createtime,article_updatetime) VALUES (?,?,?,?,?,?,?,?);`,
		param.ArticleHeader.ID,
		param.ArticleHeader.Category.ID,
		param.ArticleHeader.Title,
		1,  // 文章作者
		"", // 文章摘要
		param.Content,
		param.ArticleHeader.CreateTime,
		param.ArticleHeader.CreateTime,
	)
	if err != nil {
		log.Printf("dao.InsertArticle 报错errror:%v", err)
		return false
	} else {
		//id, _ := result.LastInsertId()
		//fmt.Println(int32(id))
	}
	return true
}

// 根据传入参数，更新对应文章
func (s *DBServer) UpdateArticle(param *definition.SaveArticleInfo) bool {
	//if
	return true
}
