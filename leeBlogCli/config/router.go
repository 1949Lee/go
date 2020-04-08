package config

var (
	// websocket编译markdown的服务地址路由。
	WebsocketParserPath = "/ws/parser"

	// 扫码登录的websocket地址路由
	WebsocketCheckLoginPath = "/ws/check-login"
	Site                    = "/"

	// 确认登录路由
	ConfirmLogin = Site + "confirm-login"

	// 接口路由 文章添加新附件
	NewFile = Site + "new-file"

	// 接口路由 删除文章附件
	DeleteFile = Site + "delete-file"

	// 接口路由 获取文章的一个新ID
	NewArticleID = Site + "new-article-id"

	// 接口路由 获取分类，并查询每个分类下的标签
	TagsGroupByCategory = Site + "categories-with-tags"

	// 接口路由 添加新标签
	NewTag = Site + "new-tag"

	// 接口路由 删除标签
	DeleteTag = Site + "delete-tag"

	// 接口路由 添加新分类
	NewCategory = Site + "new-category"

	// 接口路由 删除分类
	DeleteCategory = Site + "delete-category"

	// 接口路由 获取文章信息
	GetArticleWithEditingInfo = Site + "get-article-and-info"

	// 接口路由 新增或更新文章信息
	SaveArticle = Site + "save-article"

	// 接口路由 新增或更新文章信息
	ArticleList = Site + "article-list"

	// 接口路由 展示文章
	ShowArticle = Site + "show-article"

	// 静态资源路由
	FileResource = Site + "static/"
)
