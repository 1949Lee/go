package config

var (
	// websocket编译markdown的服务地址路由。
	WebsocketParserPath = "/ws/parser"
	Site                = "/"

	// 接口路由 文章添加新附件
	NewFile = Site + "new-file"

	// 接口路由 删除文章附件
	DeleteFile = Site + "delete-file"

	// 接口路由 获取文章的一个新ID
	NewArticleID = Site + "new-article-id"

	// 接口路由 获取分类，并查询每个分类下的标签
	TagsGroupByCategory = Site + "categories-with-tags"

	// 静态资源路由
	FileResource = Site + "static/"
)
