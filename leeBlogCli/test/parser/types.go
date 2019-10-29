package parser

// 全局变量
var (
	// 有序列表的序标的数字的间隔符。
	AutoOrderListLevelIndexDivider = "."

	// 行内状态含义列表
	LineState = LineStateEnum{
		Start:               "1",
		ItalicStart:         "2",
		ItalicEnd:           "2-1",
		DeletedTextStart:    "3",
		DeletedTextEnd:      "3-1",
		LinkTextStart:       "4",
		LinkTextEnd:         "4-1",
		LinkHrefStart:       "4-2",
		LinkHrefEnd:         "4-3",
		BackgroundStrongEnd: "5",
		ImageTextStart:      "6",
		ImageTextEnd:        "6-1",
		ImageHrefStart:      "4-2",
		ImageHrefEnd:        "4-3",
	}

	// 需要转义的字符
	markdownRunes = markdownLexicalRunes{
		backgroundRune:    '`',
		autoOrderListRune: '+',
		listRune:          '-',
		linkRune:          '[',
		imageRune:         '!',
		quoteRune:         '>',
		deleteRune:        '~',
		headerRune:        '#',
		italicRune:        '*',
	}
)

// 行内状态类型
type LineStateType string

// 行内状态枚举
type LineStateEnum struct {
	// 开始态
	Start LineStateType

	// 斜体*开始
	ItalicStart LineStateType

	// 斜体*结束
	ItalicEnd LineStateType

	// 删除线~开始
	DeletedTextStart LineStateType

	// 删除线~结束
	DeletedTextEnd LineStateType

	// 链接文案开始
	LinkTextStart LineStateType

	// 链接文案结束
	LinkTextEnd LineStateType

	// 链接URL开始
	LinkHrefStart LineStateType

	// 链接URL结束
	LinkHrefEnd LineStateType

	// 行内变色强调结束
	BackgroundStrongEnd LineStateType

	// 图像说明文案开始
	ImageTextStart LineStateType

	// 图像说明文案结束
	ImageTextEnd LineStateType

	// 图像src文案结束
	ImageHrefStart LineStateType

	// 图像src文案结束
	ImageHrefEnd LineStateType
}

// Markdown的每一行
type Line struct {
	// 原始字符串
	Origin []rune

	// 行的状态
	state LineStateType

	// 行已经确定的token
	Tokens []Token

	// 行未处理的准markdown字符数组
	unresolvedTokens []unresolvedToken

	// 行临时的字符内容的开始下标
	textStart int
}

// Markdown的每一块
type BlockResult struct {
	TokenType   string
	IndentCount int
}

// 行内未处理的准markdown字符
type unresolvedToken struct {
	text              rune
	start             bool
	tokenType         string
	contentTokenStart int
}

// Token的html节点的属性
type NodeAttr struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 需要转义的字符
type markdownLexicalRunes struct {
	backgroundRune    rune
	autoOrderListRune rune
	listRune          rune
	linkRune          rune
	imageRune         rune
	quoteRune         rune
	deleteRune        rune
	headerRune        rune
	italicRune        rune
}

// markdown的字符结构体
type Token struct {
	Text        string     `json:"text"`
	TokenType   string     `json:"tokenType"`
	NodeTagName string     `json:"tagName"`
	NodeClass   string     `json:"class"`
	NodeAttrs   []NodeAttr `json:"attrs"`
	Children    []Token    `json:"children"`
}

// Token数组
type TokenSlice []Token
