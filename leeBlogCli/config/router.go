package config

var (
	// websocket编译markdown的服务地址路由。
	WebsocketParserPath = "/ws/parser"
	Site                = "/"
	NewFile             = Site + "new-file"
	DeleteFile          = Site + "delete-file"
	FileResource        = Site + "static/"
)
