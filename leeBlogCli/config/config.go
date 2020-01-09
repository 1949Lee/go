package config

import "time"

var (
	// 一段时间内，将会发送前端最后一次请求的编译结果，防止瞬间的并发请求。这段时间默认为1500ms
	WebsocketTickerDuration time.Duration = time.Millisecond * 1500

	// 服务开启的端口号
	ServerPort string = "1314"

	// 文章所用静态资源（图片等）的存放目录
	FilePath string = "article-file/"

	// 文章编辑后保存的草稿
	DraftPath string = "editing-draft/"

	// 文章摘要的长度
	SummaryLength int = 100

	// 首页文章列表接口的页大小的默认值
	IndexArticleListPageSize int = 6
)
