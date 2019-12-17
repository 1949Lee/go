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
	Type uint8 `json:"type"`
	Time *int  `json:"-" `

	// code码开头第一位表示type类型的值。如文件上传相关则为形如：3XX
	Code     ResponseCodeType `json:"code"`
	Data     interface{}      `json:"data"`
	Files    interface{}      `json:"files"`
	Markdown interface{}      `json:"markdown"`
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
