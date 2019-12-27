package definition

type GetArticleParamArticleType int

type getArticleParamTypeEnum struct {
	// 未发布未保存的文章
	DraftArticle GetArticleParamArticleType

	// 已发布的文章
	PublicArticle GetArticleParamArticleType
}

var GetArticleParamTypeEnum = getArticleParamTypeEnum{
	DraftArticle:  1,
	PublicArticle: 2,
}

// 获取文章和文章信息的请求的参数
type GetArticleParam struct {
	Type      GetArticleParamArticleType
	ArticleID int32
}

// 新增文章或者编辑已发布的文章时需要的信息
type EditingArticleInfo struct {
}
