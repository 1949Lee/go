package config

import "time"

var (
	// 一段时间内，将会发送前端最后一次请求的编译结果，防止瞬间的并发请求。这段时间默认为200ms
	WebsocketTickerDuration time.Duration = time.Millisecond * 200

	// 服务开启的端口号
	ServerPort string = "1314"

	FilePath string = "article-file/"
)
