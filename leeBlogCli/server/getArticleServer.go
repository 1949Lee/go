package server

import "leeBlogCli/definition"

// 从数据库查询文章信息
func (b *Blog) GetArticleInfo(param *definition.GetArticleParam) {
	b.Dao.GetArticleInfo(param)
}
