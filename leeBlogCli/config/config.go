package config

import "time"

var (
	// 程序运行环境：dev表示本地，prod表示线上。
	ENV string = ""

	// API的完成URL，形如：http://blogserver.jiaxuanlee.com
	APIFullURL = ""

	LegalOriginURL = ""

	// 项目URL，
	Self_URL string = ""

	// 一段时间内，将会发送前端最后一次请求的编译结果，防止瞬间的并发请求。这段时间默认为3000ms
	WebsocketTickerDuration time.Duration = time.Millisecond * 3000

	// 服务开启的端口号
	ServerPort string = "1016"

	StaticPath string = "lee-blog/"

	// 文件显示的路径
	FileSourcePath string = "article-file/"

	// 文章所用静态资源（图片等）的存放目录
	FilePath string = StaticPath + "article-file/"

	// 文章编辑后保存的草稿
	DraftPath string = StaticPath + "editing-draft/"

	// 文章摘要的长度
	SummaryLength int = 100

	// 首页文章列表接口的页大小的默认值
	IndexArticleListPageSize int = 6

	ImageTypeNeedConvert string = ".jpg,.jpeg,.png,.svg,.gif"
)

// 配置初始化
func Init() {
	switch ENV {
	case "dev":
		Self_URL = "localhost"
		APIFullURL = "http://" + Self_URL + ":" + ServerPort
		LegalOriginURL = "localhost:8080"
	case "prod":
		Self_URL = "blogserver.jiaxuanlee.com"
		APIFullURL = "https://" + Self_URL
		LegalOriginURL = "jiaxuanlee.com"
	}
}
