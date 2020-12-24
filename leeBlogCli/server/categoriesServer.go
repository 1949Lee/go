package server

import "leeBlogCli/definition"

// 获取所有分类和分类下的所有标签，每个标签附带含有标签的文章ID数组
func (b *Blog) GetTagsWithArticleID() (c definition.CategoryWithArticleIDListResult) {
	var err error
	c, err = b.Dao.SelectTagsWithArticleID()
	if err == nil {
		return c
	} else {
		c = definition.CategoryWithArticleIDListResult{
			List: make([]definition.CategoryWithTagsAndArticleID, 0),
		}
		return c
	}
}

// 根据文章ID获取文章列表
func (b *Blog) GetArticleListByID(param *definition.ArticleListByIDParam) (c definition.ArticleListResult) {
	return b.Dao.GetArticleListByID(param)
}
