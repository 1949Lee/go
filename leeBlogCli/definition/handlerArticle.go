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
	Type      GetArticleParamArticleType `json:"type"`
	ArticleID int32                      `json:"articleID"`
}

// 新增文章或者编辑已发布的文章时需要的信息
type EditingArticleInfo struct {
	// markdown编辑字符串
	Markdown string `json:"markdown"`

	// 文章信息
	ArticleHeader ArticleHeader `json:"articleHeader"`
}

// 文章信息
type ArticleHeader struct {
	// 文章标题
	Title string `json:"title"`

	// 文章ID
	ID int32 `json:"articleID"`

	// 文章创建日期
	CreateTime string `json:"createTime"`

	// 文章更新时间
	UpdateTime string `json:"updateTime"`

	// 文章分类
	Category Category `json:"category"`

	// 文章标签
	Tags []Tag `json:"tags"`
}

// 新增或跟新时传入的参数
type SaveArticleInfo struct {

	// 文章的类型，
	Type GetArticleParamArticleType `json:"type"`

	// 文章的markdown内容
	Content string `json:"content"`

	// 文章信息
	ArticleHeader ArticleHeader `json:"info"`
}
