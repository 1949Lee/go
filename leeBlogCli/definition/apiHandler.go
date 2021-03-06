package definition

type ParamEditingArticle struct {
	ArticleID int
	Type      uint8
	Text      string
	Files     []File
}

type ResponseCodeType int

type ResponseResult struct {
	// 1表示markdown相关，2 表示文件准备相关，3表示文件上传相关。4表示http删除上传文件相关。
	Type uint8  `json:"type"`
	Time *int64 `json:"-" `

	// code码开头第一位表示type类型的值。如文件上传相关则为形如：3XX
	Code     ResponseCodeType `json:"code"`
	Data     interface{}      `json:"data"`
	Files    interface{}      `json:"files"`
	Markdown interface{}      `json:"markdown"`
}

type APIResult struct {
	// code码开头第一位表示type类型的值。如文件上传相关则为形如：3XX
	Code    ResponseCodeType `json:"code"`
	Data    interface{}      `json:"data"`
	Message interface{}      `json:"message"`
}

type ResponseResultQueue []*ResponseResult

func (q *ResponseResultQueue) Max() *ResponseResult {
	list := *q
	var result *ResponseResult
	if len(list) > 0 {
		for i := len(list) - 1; i >= 0; i-- {
			if result == nil {
				result = list[i]
			}
			if *(result.Time) < *(list[i].Time) {
				result = list[i]
			}
		}
	}
	return result
}

type FileResponseItem struct {
	// 文件名
	FileName string `json:"fileName"`

	// 文件URL
	URL string `json:"url"`
}

// 服务器收到文件后，返回的结果体
type FileResponse []FileResponseItem

// 手机端确认登录接口参数
type ConfirmLoginParam struct {
	// 用户邮箱
	Email string `json:"key"`

	// email和设备的uuid拼接后的md5串
	Passport string `json:"passport"`
}

// 登录二维码的key
type GetLoginKeyResponse struct {
	Email string `json:"key"`
}

// 扫码登录成功后的返回参数
type ConfirmLoginResponse struct {

	// 登录的token
	LeeToken string `json:"leeToken"`
}

// 返回码的枚举
var ResponseServerCode = struct {
	// 未登录
	NotLogin ResponseCodeType
}{
	NotLogin: 1479,
}

type InterceptorOptions struct {
	// 是否需要登录，true，表示需要；false，表示不需要。
	CheckLogin bool
}
