package definition

import "leeBlogCli/parser"

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

	// 文章内容转换为html的串。
	Text string `json:"text"`
}

// 文章列表接口的参数
type ArticleListParam struct {
	// 页大小
	PageSize int `json:"pageSize"`

	// 页码
	PageIndex int `json:"pageIndex"`

	// 类别ID
	CategoryID int32 `json:"categoryID"`

	// 标签ID
	TagIDs string `json:"tagIDs"`

	// 文章标题
	Title string `json:"title"`
}

// 文章列表的一项
type ArticleListResultItem struct {
	// 文章
	Article

	// 文章类别
	CategoryName string `json:"categoryName" db:"ctg_name"`

	// 文章标签
	Tags []Tag `json:"tags"`
	//Tag
}

// 文章列表的一项
type ArticleListResult struct {
	// 文章列表
	List []ArticleListResultItem `json:"list"`

	// 是否是最后一页
	IsLastPage bool `json:"isLastPage"`
}

// 展示文章入参
type ShowArticleParam struct {
	ID int32 `json:"id"`
}

// 展示文章入参
type ShowArticleResult struct {
	List    []parser.TokenSlice   `json:"list"`
	Text    string                `json:"text"`
	Html    string                `json:"html"`
	Article ArticleListResultItem `json:"article"`
}
