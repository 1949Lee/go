package server

import (
	"leeBlogCli/definition"
	"leeBlogCli/parser"
)

// 从数据库查询文章信息
func (b *Blog) GetArticleHeader(param *definition.GetArticleParam) definition.Article {
	return b.GetArticle(param)
}

// 从数据库查询文章信息
func (b *Blog) SaveArticle(param *definition.SaveArticleInfo) bool {
	if param.Type == definition.GetArticleParamTypeEnum.DraftArticle { // 新增文章走添加接口
		return b.Dao.InsertArticle(param)
	} else if param.Type == definition.GetArticleParamTypeEnum.PublicArticle { // 更新文章走更新接口
		return b.Dao.UpdateArticle(param)
	}
	return false
}

// 获取文章的全部信息
func (b *Blog) GetArticle(param *definition.GetArticleParam) definition.Article {
	return b.Dao.GetArticle(param.ArticleID)
}

// 根据传入的查询条件（分页、关键字、分类、标签等）从数据库查询文章列表
func (b *Blog) ArticleList(param *definition.ArticleListParam) definition.ArticleListResult {
	return b.Dao.GetArticleList(param)
}

// 根据传入的文章id从数据库获取文章内容，然后经过parser的转换后，转换成HTMl的串。
func (b *Blog) ShowArticle(param *definition.ShowArticleParam) definition.ShowArticleResult {
	article := b.Dao.GetArticle(param.ID)
	if article.ID == 0 {
		return definition.ShowArticleResult{}
	}
	category, tags := b.Dao.GetArticleCategoryAndTags(param.ID)
	result := MarkdownParse(article.Content)
	result.Article = definition.ArticleListResultItem{
		Article:      article,
		CategoryName: category.Name,
		Tags:         tags,
	}
	result.Article.Content = ""
	return result
}

func MarkdownParse(p string) definition.ShowArticleResult {
	dataList, _ := parser.MarkdownParse(p)
	return definition.ShowArticleResult{
		Text: p,
		List: dataList,
	}
}
