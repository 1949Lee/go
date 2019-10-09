package parser

import (
	"bytes"
	"strconv"
	"strings"
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

	// 链接URL结束
	BackgroundStrongEnd LineStateType
}

// 行内状态含义列表
var LineState = LineStateEnum{
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

// 未处理的准markdown字符
type unresolvedToken struct {
	text              rune
	start             bool
	tokenType         string
	contentTokenStart int
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

// html节点的属性
type NodeAttr struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 根据当前line的Tokens生成html
func (l *Line) ToHtml() string {
	return lineToHtml(l.Tokens)
}

// 行内转换的解析函数
func (l *Line) LineParse() {
	if len(l.Origin) == 0 {
		l.Tokens = append(l.Tokens, Token{TokenType: "empty-line-br", NodeTagName: "br", NodeClass: "empty-line-br"})
	} else {
		//l.ItalicTextParse()
		//l.DeleteTextParse()
		//l.LinkTextParse()
		//l.BackgroundStrongParse()
		l.HeaderTitleParse()
	}
}

// 斜体、粗体、粗斜体转换方法
func (l *Line) ItalicTextParse() {
	l.state = LineState.Start
	l.textStart = -1
	for i := 0; i < len(l.Origin); i++ {
		ch := l.Origin[i]
		if ch == '\\' && i < len(l.Origin)-1 && (l.Origin[i+1] == '*') {
			l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
		}
		switch l.state {
		case LineState.Start:
			l.unresolvedTokens = []unresolvedToken{}
			switch ch {
			case '*':

				// 1. 遇到*之后。需要记录这个*留着判断。
				ut := unresolvedToken{text: '*', start: true, contentTokenStart: len(l.Tokens)}
				l.unresolvedTokens = append(l.unresolvedTokens, ut)
				l.state = LineState.ItalicStart

				// 2. 并且还要把*之前的token的text（若有）。
				if l.textStart != -1 {
					l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
					l.textStart = -1
					l.unresolvedTokens[len(l.unresolvedTokens)-1].contentTokenStart = len(l.Tokens)
				}

				continue
			default:
				// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
				}
				continue
			}
		case LineState.ItalicStart:
			switch ch {
			case '*':
				if l.textStart == -1 {
					l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '*', start: true})
				}
				continue
			default:
				// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
					l.state = LineState.ItalicEnd
				}
				continue
			}
		case LineState.ItalicEnd:
			switch ch {
			case '*':
				length := len(l.unresolvedTokens)
				if length > 0 {
					i = l.confirmItalicType(i)
					l.state = LineState.Start
				} else {
					l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '*', start: true})

					if l.textStart != -1 {
						l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
						l.textStart = -1
					}
					l.state = LineState.ItalicStart
				}
				continue
			default:
				// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
				}
				continue
			}
		}
	}
	l.resolveLineToken()
}

// 删除线转换方法，会自行调用ItalicTextParse（斜体、粗体、粗斜体转换方法）
func (l *Line) DeleteTextParse() {
	l.state = LineState.Start
	l.textStart = -1
	for i := 0; i < len(l.Origin); i++ {
		ch := l.Origin[i]
		if ch == '\\' && i < len(l.Origin)-1 && (l.Origin[i+1] == '~') {
			l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
		}
		switch l.state {
		case LineState.Start:
			l.unresolvedTokens = []unresolvedToken{}
			switch ch {
			case '~':

				// 1. 遇到*之后。需要记录这个*留着判断。
				l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '~', start: true, contentTokenStart: len(l.Tokens)})
				l.state = LineState.DeletedTextStart

				// 2. 并且还要把*之前的token的text（若有）。
				if l.textStart != -1 {
					l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
					l.textStart = -1
					l.unresolvedTokens[len(l.unresolvedTokens)-1].contentTokenStart = len(l.Tokens)
				}

				continue
			default:
				// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
				}
				continue
			}
		case LineState.DeletedTextStart:
			switch ch {
			case '~':
				if l.textStart == -1 {
					if len(l.unresolvedTokens) == 2 {
						l.textStart = i
						l.state = LineState.DeletedTextEnd
					} else {
						l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '~', start: true})
					}
				}
				continue
			default:
				// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					if len(l.unresolvedTokens) == 2 {
						l.textStart = i
						l.state = LineState.DeletedTextEnd
					} else {
						l.textStart = i - 1
						l.unresolvedTokens = l.unresolvedTokens[:len(l.unresolvedTokens)-1]
						l.state = LineState.Start
					}
				} else {
					if len(l.unresolvedTokens) == 2 {
						l.state = LineState.DeletedTextEnd
					} else {
						l.unresolvedTokens = l.unresolvedTokens[:len(l.unresolvedTokens)-1]
						l.state = LineState.Start
					}
				}
				continue
			}
		case LineState.DeletedTextEnd:
			switch ch {
			case '~':
				length := len(l.unresolvedTokens)
				if length > 0 {
					i = l.confirmDeletedText(i)
					l.state = LineState.Start
				} else {
					l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '~', start: true})
					if l.textStart != -1 {
						l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
						l.textStart = -1
					}
					l.state = LineState.DeletedTextStart
				}
				continue
			default:
				// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
				}
				continue
			}
		}

	}
	l.resolveLineToken()
	l.parseWithOther(func(line *Line) {
		line.ItalicTextParse()
	})
}

// 链接转换方法，会自行调用DeleteTextParse
func (l *Line) LinkTextParse() {
	//linkRegex := regexp.MustCompile(`\[[^]]+]\([^)]+\)`)
	//links := linkRegex.FindAllStringIndex(string(l.Origin),-1)
	//fmt.Println(links)
	l.state = LineState.Start
	l.textStart = -1
	tempToken := Token{}
	l.unresolvedTokens = []unresolvedToken{}
	for i := 0; i < len(l.Origin); i++ {
		ch := l.Origin[i]
		if ch == '\\' && i < len(l.Origin)-1 && (l.Origin[i+1] == '[' ||
			l.Origin[i+1] == ']' ||
			l.Origin[i+1] == ')' ||
			l.Origin[i+1] == '(') {
			l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
		}
		switch l.state {
		case LineState.Start:
			switch ch {
			case '[':
				// 1. 遇到[之后。需要记录这个[留着判断。
				l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '[', start: true})
				l.state = LineState.LinkTextEnd

				// 2. 并且还要把[之前的token的text（若有）。
				if l.textStart != -1 {
					l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
					l.textStart = -1
				}

				continue
			default:
				// 如果读到非语法字符,则记录位置。否则继续扫描下一个字符。textStart == -1表示只记录最开始的位置
				if l.textStart == -1 {
					l.textStart = i
				}
				continue
			}
		case LineState.LinkTextEnd:
			switch ch {
			case ']':
				if len(l.unresolvedTokens) == 1 {
					// 如果]中有文案，且后面就是(则认为已经进入链接URL判断的状态
					if l.textStart != -1 && i+1 < len(l.Origin) && l.Origin[i+1] == '(' {
						tempToken.Text = string(l.Origin[l.textStart:i])
						tempToken.TokenType = "web-link"
						l.unresolvedTokens = l.unresolvedTokens[:len(l.unresolvedTokens)-1]
						l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '(', start: true})
						l.state = LineState.LinkHrefEnd
						i++
						l.textStart = -1
					} else if l.textStart != -1 && i+1 < len(l.Origin) && l.Origin[i+1] != '(' {
						l.appendNewToken(&Token{Text: "[" + string(l.Origin[l.textStart:i]) + "]", TokenType: "text"})
						l.state = LineState.Start
						l.textStart = -1
						tempToken = Token{}
						l.unresolvedTokens = []unresolvedToken{}
						l.textStart = -1
					} else {
						// 如果已经读到了行的最后一个字符，则进行一些未完成的token处理
						if i+1 == len(l.Origin) {
							if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
								l.Tokens[len(l.Tokens)-1].Text += "[" + string(l.Origin[l.textStart:]) + "]"
								updateToken(&l.Tokens[len(l.Tokens)-1])
							} else {
								l.appendNewToken(&Token{Text: "[" + string(l.Origin[l.textStart:]) + "]", TokenType: "text"})
							}
							l.textStart = -1
							l.unresolvedTokens = []unresolvedToken{}
						}
						l.state = LineState.Start
						l.textStart = -1
						tempToken = Token{}
						l.unresolvedTokens = []unresolvedToken{}
					}
				}
				continue
			default:
				// 如果读到非语法字符，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
				}

				// 如果已经读到了行的最后一个字符，则进行一些未完成的token处理
				if i+1 == len(l.Origin) {
					if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
						l.Tokens[len(l.Tokens)-1].Text += "[" + string(l.Origin[l.textStart:])
						updateToken(&l.Tokens[len(l.Tokens)-1])
					} else {
						l.appendNewToken(&Token{Text: "[" + string(l.Origin[l.textStart:]), TokenType: "text"})
					}
					l.textStart = -1
					l.unresolvedTokens = []unresolvedToken{}
				}
				continue
			}
		case LineState.LinkHrefEnd:
			switch ch {
			case ')':
				length := len(l.unresolvedTokens)
				//appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
				if length == 1 {
					if l.textStart != -1 {
						tempToken.NodeAttrs = []NodeAttr{
							{Key: "href", Value: string(l.Origin[l.textStart:i])},
						}
						l.appendNewToken(&tempToken)
					} else {
						var builder strings.Builder
						builder.WriteString("[")
						builder.WriteString(tempToken.Text)
						builder.WriteString("]()")
						l.appendNewToken(&Token{Text: builder.String(), TokenType: "text"})
					}
				}
				l.state = LineState.Start
				l.textStart = -1
				tempToken = Token{}
				l.unresolvedTokens = []unresolvedToken{}
				continue
			default:
				// 如果读到非语法字符，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
				}
				// 如果已经读到了行的最后一个字符，则进行一些未完成的token处理
				if i+1 == len(l.Origin) {
					if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
						l.Tokens[len(l.Tokens)-1].Text += "[" + tempToken.Text + "](" + string(l.Origin[l.textStart:])
						updateToken(&l.Tokens[len(l.Tokens)-1])
					} else {
						l.appendNewToken(&Token{Text: "[" + tempToken.Text + "](" + string(l.Origin[l.textStart:]), TokenType: "text"})
					}
					l.textStart = -1
					l.unresolvedTokens = []unresolvedToken{}
				}
				continue
			}
		}

	}
	l.resolveLineToken()
	l.parseWithOther(func(line *Line) {
		line.DeleteTextParse()
	})
}

//行内底色变色强调的转换方法，会自行调用LinkTextParse
func (l *Line) BackgroundStrongParse() {
	l.state = LineState.Start
	l.textStart = -1
	for i := 0; i < len(l.Origin); i++ {
		ch := l.Origin[i]
		if ch == '\\' && i < len(l.Origin)-1 && (l.Origin[i+1] == '`') {
			l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
		}
		switch l.state {
		case LineState.Start:
			l.unresolvedTokens = []unresolvedToken{}
			switch ch {
			case '`':
				l.state = LineState.BackgroundStrongEnd

				// 2. 并且还要把*之前的token的text（若有）。
				if l.textStart != -1 {
					l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
					l.textStart = -1
				}

				continue
			default:
				// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
				}
				continue
			}
		case LineState.BackgroundStrongEnd:
			switch ch {
			case '`':
				if l.textStart != -1 {
					text := string(l.Origin[l.textStart:i])
					l.appendNewToken(&Token{Text: text, TokenType: "background-strong"})
					l.state = LineState.Start
					l.unresolvedTokens = []unresolvedToken{}
					l.textStart = -1
				}
				continue
			default:
				// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
				}
				continue
			}
		}

	}
	l.resolveLineToken()
	l.parseWithOther(func(line *Line) {
		line.LinkTextParse()
	})
}

//多级标题转换方法，会自行调用BackgroundStrongParse
func (l *Line) HeaderTitleParse() {
	l.state = LineState.Start
	l.textStart = -1
	//如果行开头是#号，则进行多级标题判断。
	if l.Origin[0] == '#' {
		inHeader := true
		headerLevel := 0
		for i := 0; i < len(l.Origin); i++ {
			ch := l.Origin[i]
			if inHeader && ch == '#' {
				headerLevel++
			} else {
				if inHeader {
					inHeader = false
					if ch == ' ' { // 连续#号之后，必须跟一个空格。
						l.textStart = i + 1
					} else { // 连续#号之后，若没有跟空格，则直接进行其他转换。
						l.BackgroundStrongParse()
						return
					}
				}
				if ch == '\\' && i < len(l.Origin)-1 && (l.Origin[i+1] == '#') {
					l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
				}
			}
		}
		l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:]), NodeTagName: "h" + strconv.Itoa(headerLevel), TokenType: "header"})
		l.parseWithOther(func(line *Line) {
			line.BackgroundStrongParse()
		})
	} else { // 行开头不是#，直接进行其他转换
		for i := 0; i < len(l.Origin); i++ {
			ch := l.Origin[i]
			if ch == '\\' && i < len(l.Origin)-1 && (l.Origin[i+1] == '#') {
				l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
			}
		}
		l.BackgroundStrongParse()
	}

}

// 行内判断斜体的具体token的类型的结束函数，token具体是斜体、加粗、粗斜体
func (l *Line) confirmItalicType(i int) int {
	// 保留初始流的读取位置下标
	originIndex := i
	var tokenType string

	//之前l.unresolvedTokens中存储了连续的'开始*'，所以现在需要读取i之后的连续'结束*'。
	//状态机处于State11时，读到*才调用此函数。所以默认已经读取了一个结束*
	endCount := 1 // endCount表示token的有效*的个数。前后各一个*表示斜体，前后各两个*表示粗体，前后各三个*表示粗斜体。
	i++
	l.unresolvedTokens = l.unresolvedTokens[:len(l.unresolvedTokens)-1]

	// for循环读取i之后的连续'结束*'。
	for ; i < len(l.Origin); i++ {
		// 如果读取的下一个字符是*，并且。可以与l.unresolvedTokens中的'开始*'匹配。则，有效*的数目+1
		if l.Origin[i] == '*' && len(l.unresolvedTokens) > 0 {
			l.unresolvedTokens = l.unresolvedTokens[:len(l.unresolvedTokens)-1]
			endCount++
		} else { // 不符合条件时，往前一个字符。不影响后续读取字符。
			i--
			break
		}
	}

	//根据有效*的数目确定类型。前后各一个*表示斜体，前后各两个*表示粗体，前后各三个*表示粗斜体。
	switch endCount {
	case 1: // 斜体
		tokenType = "italic"
	case 2: // 粗体
		tokenType = "bold"
	case 3: // 粗斜体
		tokenType = "bold-italic"
	}

	// 返回正确的token内容
	if l.textStart != -1 {
		var temp string

		if len(l.unresolvedTokens) > 0 {

			//遗留的'开始*'需要放到内容的前方
			temp = joinTokens(l.unresolvedTokens, "") + string(l.Origin[l.textStart:originIndex])

			//制空开始数组。
			l.unresolvedTokens = []unresolvedToken{}
		} else {
			temp = string(l.Origin[l.textStart:originIndex])
		}
		l.appendNewToken(&Token{Text: temp, TokenType: tokenType})
		l.textStart = -1
	} else {
		l.appendNewToken(&Token{Text: "", TokenType: tokenType})
	}
	return i
}

// 行内判断删除线token的结束函数
func (l *Line) confirmDeletedText(i int) int {
	// 保留初始流的读取位置下标
	originIndex := i
	var tokenType string

	endCount := 1 // endCount表示token的有效~的个数。前后各2个~表示删除线。
	i++
	l.unresolvedTokens = l.unresolvedTokens[:len(l.unresolvedTokens)-1]

	// 如果读取的下一个字符是*，并且。可以与l.unresolvedTokens中的'开始*'匹配。则，有效*的数目+1
	if i < len(l.Origin) && l.Origin[i] == '~' && len(l.unresolvedTokens) > 0 {
		l.unresolvedTokens = l.unresolvedTokens[:len(l.unresolvedTokens)-1]
		endCount++
	} else { // 不符合条件时，往前一个字符。不影响后续读取字符。
		i--
	}

	//根据有效~的数目确定类型。前后各2个~表示删除线。
	if endCount == 2 { // 删除线成立
		tokenType = "deleted-text"
	} else {
		l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '~', start: true})
		return i
	}

	// 返回正确的token内容
	if l.textStart != -1 {
		var temp string
		//
		if len(l.unresolvedTokens) > 0 {
			//
			//    //遗留的'开始*'需要放到内容的前方
			//    temp = joinTokens(l.unresolvedTokens, "") + string(l.Origin[l.textStart:originIndex])
			//
			//    //制空开始数组。
			//   l.unresolvedTokens = []unresolvedToken{}
		} else {
			//    temp = string(l.Origin[l.textStart:originIndex])
		}
		temp = string(l.Origin[l.textStart:originIndex])
		l.appendNewToken(&Token{Text: temp, TokenType: tokenType})
		l.textStart = -1
	} else {
		l.appendNewToken(&Token{Text: "", TokenType: tokenType})
	}
	return i
}

// 遍历完成之后将行进行最后的结尾工作：处理未解决的token、处理未发射为text的字符
func (l *Line) resolveLineToken() {
	if len(l.unresolvedTokens) > 0 {
		//l.Tokens = append(l.Tokens, "*")
		if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
			l.Tokens[len(l.Tokens)-1].Text += joinTokens(l.unresolvedTokens, "")
			updateToken(&l.Tokens[len(l.Tokens)-1])
		} else {
			l.appendNewToken(&Token{Text: joinTokens(l.unresolvedTokens, ""), TokenType: "text"})
		}
	}
	if l.textStart != -1 {
		if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
			l.Tokens[len(l.Tokens)-1].Text += string(l.Origin[l.textStart:])
			updateToken(&l.Tokens[len(l.Tokens)-1])
		} else {
			l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:]), TokenType: "text"})
		}
	}
}

// 向line中添加新的token
func (l *Line) appendNewToken(t *Token) {
	updateToken(t)
	l.Tokens = append(l.Tokens, *t)
}

// 将当前行的Tokens进行其他Markdown语法的二次转换、转换的方法为参数。转换后，当前行的Tokens将会更新。
func (l *Line) parseWithOther(parseFunc func(*Line)) {
	tokens := make([]Token, 0)
	for i := range l.Tokens {
		lineText := Line{Origin: []rune(l.Tokens[i].Text), Tokens: []Token{}}
		parseFunc(&lineText)
		tokens = append(tokens, l.Tokens[i].updateWith(lineText.Tokens...)...)
	}
	l.Tokens = tokens
}

// 用新的token去更新某个token
func (t Token) updateWith(tokens ...Token) []Token {
	if t.TokenType == "text" {
		return tokens
	}
	if len(tokens) == 1 && t.NodeTagName == tokens[0].NodeTagName {
		if tokens[0].NodeClass == "text" {
			tokens[0].NodeClass = t.NodeClass
		} else {
			tokens[0].NodeClass += " " + t.NodeClass
		}
		if len(t.NodeAttrs) > 0 {
			tokens[0].NodeAttrs = append(tokens[0].NodeAttrs, t.NodeAttrs...)
		}
		return tokens
	}
	switch t.TokenType {
	case "deleted-text":
		for i := range tokens {
			if tokens[i].NodeClass == "text" {
				tokens[i].NodeClass = t.NodeClass
			} else {
				tokens[i].NodeClass += " " + t.NodeClass
			}
			if len(t.NodeAttrs) > 0 {
				tokens[i].NodeAttrs = append(tokens[i].NodeAttrs, t.NodeAttrs...)
			}
		}
		return tokens
	case "web-link", "background-strong", "header":
		t.Children = tokens
		t.Text = ""
		return []Token{t}
	}
	return tokens
}

// 将传入的token数组转化为字符串，数组中的每一项转化为字符串之后用str连接起来形成一个新的字符串
func joinTokens(p []unresolvedToken, str string) string {
	result := ""
	var buffer bytes.Buffer
	for _, t := range p {
		buffer.WriteString(string(t.text))
		if str != "" {
			buffer.WriteString(str)
		}
	}
	result = buffer.String()

	return result
}

// 根据传入的token的不同类型来做相应的处理
func updateToken(t *Token) {
	switch t.TokenType {
	case "text":
		t.NodeClass = "text"
		t.NodeTagName = "span"
	case "italic":
		t.NodeClass = "italic"
		t.NodeTagName = "span"
	case "bold":
		t.NodeClass = "bold"
		t.NodeTagName = "span"
	case "bold-italic":
		t.NodeClass = "bold-italic"
		t.NodeTagName = "span"
	case "deleted-text":
		t.NodeClass = "deleted-text"
		t.NodeTagName = "span"
	case "web-link":
		t.NodeClass = "inline-web-link"
		t.NodeTagName = "a"
	case "background-strong":
		t.NodeClass = "inline-background-strong"
		t.NodeTagName = "span"
	case "header":
		t.NodeClass = "header-" + t.NodeTagName
	}
}

// 将传入的行转换为html
func LinesToHtml(lines [][]Token) string {
	var builder strings.Builder
	builder.WriteString(`<div class="content">`)
	for i := range lines {
		builder.WriteString(lineToHtml(lines[i]))
	}
	builder.WriteString(`</div>`)
	return builder.String()
}

// 将传入的token数组转化为html
func lineToHtml(tokens []Token) string {
	var builder strings.Builder
	if tokens[0].TokenType == "empty-line-br" {
		builder.WriteString(`<div class="block empty-single-line-block">`)
		builder.WriteString(tokensToHtml(tokens))
		builder.WriteString(`</div>`)
	} else {
		builder.WriteString(`<div class="block">`)
		builder.WriteString(tokensToHtml(tokens))
		builder.WriteString(`</div>`)
	}
	return builder.String()
}

// 将token转化为html
func tokensToHtml(tokens []Token) string {
	var builder strings.Builder
	for i := range tokens {
		if tokens[i].TokenType == "empty-line-br" {
			builder.WriteString("<")
			builder.WriteString(tokens[i].NodeTagName)
			builder.WriteString(` class="`)
			builder.WriteString(tokens[i].NodeClass)
			builder.WriteString(`" `)
			if len(tokens[i].NodeAttrs) > 0 {
				for j := range tokens[i].NodeAttrs {
					builder.WriteString(` `)
					builder.WriteString(tokens[i].NodeAttrs[j].Key)
					builder.WriteString(`="`)
					builder.WriteString(tokens[i].NodeAttrs[j].Value)
					builder.WriteString(`" `)
				}
			}
			builder.WriteString("/>")
		} else {
			builder.WriteString("<")
			builder.WriteString(tokens[i].NodeTagName)
			builder.WriteString(` class="`)
			builder.WriteString(tokens[i].NodeClass)
			builder.WriteString(`" `)
			if len(tokens[i].NodeAttrs) > 0 {
				for j := range tokens[i].NodeAttrs {
					builder.WriteString(` `)
					builder.WriteString(tokens[i].NodeAttrs[j].Key)
					builder.WriteString(`="`)
					builder.WriteString(tokens[i].NodeAttrs[j].Value)
					builder.WriteString(`" `)
				}
			}
			builder.WriteString(">")
			builder.WriteString(tokens[i].Text)
			if len(tokens[i].Children) > 0 {
				builder.WriteString(tokensToHtml(tokens[i].Children))
			}
			//result += "<"+ l.t.NodeTagName +" class="++" >"
			builder.WriteString("</")
			builder.WriteString(tokens[i].NodeTagName)
			builder.WriteString(">")
		}
	}
	return builder.String()
}

// 处理行line是否是多行块的开头
func handleBlockForLines(line Line, index int) {
	//l.state = LineState.Start
	//l.textStart = -1
	////如果行开头是#号，则进行多级标题判断。
	//if l.Origin[0] == '#' {
	//    inHeader := true
	//    headerLevel := 0
	//    for i := 0; i < len(l.Origin); i++ {
	//        ch := l.Origin[i]
	//        if inHeader && ch == '#' {
	//            headerLevel++
	//        } else {
	//            if inHeader {
	//                inHeader = false
	//                if ch == ' ' { // 连续#号之后，必须跟一个空格。
	//                    l.textStart = i + 1
	//                } else { // 连续#号之后，若没有跟空格，则直接进行其他转换。
	//                    l.BackgroundStrongParse()
	//                    return
	//                }
	//            }
	//            if ch == '\\' && i < len(l.Origin)-1 && (l.Origin[i+1] == '#') {
	//                l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
	//            }
	//        }
	//    }
	//    l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:]), NodeTagName: "h" + strconv.Itoa(headerLevel), TokenType: "header"})
	//    l.parseWithOther(func(line *Line) {
	//        line.BackgroundStrongParse()
	//    })
	//} else { // 行开头不是#，直接进行其他转换
	//    for i := 0; i < len(l.Origin); i++ {
	//        ch := l.Origin[i]
	//        if ch == '\\' && i < len(l.Origin)-1 && (l.Origin[i+1] == '#') {
	//            l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
	//        }
	//    }
	//    l.BackgroundStrongParse()
	//}
}

type BlockResult struct {
	TokenType   string
	IndentCount int
}

func getIndentCount(lineText string) (int, string) {
	indentCount := 0
	realStr := strings.TrimLeftFunc(lineText, func(r rune) bool {
		if r == ' ' {
			indentCount++
		}
		return r == ' '
	})
	return indentCount, realStr
}

// 判断某一行的内容是否为块开头
func isInBlock(lineText string) (bool, BlockResult) {
	//origin := []rune(lineText)
	indentCount, realRune := getIndentCount(lineText)
	if lineText == "" {
		return false, BlockResult{}
	}
	switch realRune[0] {
	case '*', '-': // 无序列表list
		if realRune[1] == ' ' {
			return true, BlockResult{TokenType: "list", IndentCount: indentCount}
		}
	case '>': // 颜色块colored-block或文字引用block-quote，需要进一步判断
		if realRune[1] == '>' && realRune[2] == '>' {
			return true, BlockResult{TokenType: "colored-block"}
		} else {
			return true, BlockResult{TokenType: "block-quote"}
		}
	case '|': // 表格table
	// TODO 表格转换待实现，这里需要做判断是否是表格
	case '`': // 代码块code-block，需要进一步判断
		if realRune[1] == '`' && realRune[2] == '`' {
			return true, BlockResult{TokenType: "code-block"}
		}
	case '+': // 自动有序列表auto-order-list
		if realRune[1] == ' ' {
			return true, BlockResult{TokenType: "auto-order-list", IndentCount: indentCount}
		}
	case ':': // 名词定义列表word-list，需要进一步判断
		if realRune[1] == ':' {
			return true, BlockResult{TokenType: "word-list"}
		}
	}
	return false, BlockResult{}
}

// 判断某一行的内容是否为属于无序列表list的一部分
func isInList(lineText string) (bool, BlockResult) {
	indentCount, realRune := getIndentCount(lineText)
	switch realRune[0] {
	case '*', '-': // 无序列表list
		if realRune[1] == ' ' {
			return true, BlockResult{TokenType: "list", IndentCount: indentCount}
		}
	}
	return false, BlockResult{}
}

// 接受markdown字符串，并将之转化为html
func MarkdownParse(markdownText string) ([][]Token, string) {
	//scanner := bufio.NewScanner(strings.NewReader(markdownText))
	//dataList := make([][]Token, 0)
	//for scanner.Scan() {
	//	lineText := scanner.Text()
	//	line := Line{Origin: []rune(lineText), Tokens: []Token{}}
	//	line.LineParse()
	//	dataList = append(dataList, line.Tokens)
	//}

	//这种split的方法比bufio那种读取块100-500微秒。
	list := strings.Split(markdownText, "\n")
	dataList := make([][]Token, 0)
	for i := 0; i < len(list); i++ {
		if ok, blockResult := isInBlock(list[i]); ok {
			switch blockResult.TokenType {
			case "list":
				var tokens []Token
				i, tokens = listParse(list, i, blockResult)
				i--
				dataList = append(dataList, tokens)
			}
		} else {
			line := Line{Origin: []rune(list[i]), Tokens: []Token{}}
			line.LineParse()
			dataList = append(dataList, line.Tokens)
		}
	}
	return [][]Token{}, LinesToHtml(dataList)
}

func listParse(lines []string, index int, blockResult BlockResult) (int, []Token) {
	tokens := []Token{
		{
			TokenType:   "list",
			NodeTagName: "ul",
			NodeClass:   "list",
			Children:    []Token{},
		}}
	originIndent := blockResult.IndentCount
	i := index
	for ; i < len(lines); i++ {
		ok, temResult := isInList(lines[i])
		if ok { // list的新一行
			if temResult.IndentCount >= originIndent && temResult.IndentCount < originIndent+4 { // 新的列表项不是子列表
				tokens[0].Children = append(tokens[0].Children, Token{
					TokenType:   "list-item",
					NodeTagName: "li",
					NodeClass:   "list-item",
					Children:    []Token{},
				})
				text := []rune(lines[i])[2+temResult.IndentCount:]
				line := Line{Origin: text, Tokens: []Token{}}
				line.LineParse()
				tokens[0].Children[len(tokens[0].Children)-1].Children = line.Tokens
			} else if temResult.IndentCount >= originIndent+4 { //新的列表项，缩进符合下一级。
				temIndex, subTokens := listParse(lines, i, temResult)
				i = temIndex - 1
				tokens[0].Children[len(tokens[0].Children)-1].Children = append(tokens[0].Children[len(tokens[0].Children)-1].Children, subTokens...)
			} else if temResult.IndentCount < originIndent {
				return i, tokens
			}
		} else { // list结束
			index = i
			return index, tokens
		}
	}
	index = i
	return index, tokens
}
