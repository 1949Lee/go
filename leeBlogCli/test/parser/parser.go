package parser

import (
	"bytes"
	"strings"
)

// 行内状态类型
type LineStateType int

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
}

// 行内状态含义列表
var LineState = LineStateEnum{
	Start:            1,
	ItalicStart:      2,
	ItalicEnd:        21,
	DeletedTextStart: 3,
	DeletedTextEnd:   31,
}

// Markdown的每一行
type Line struct {
	Origin           []rune
	state            LineStateType
	Tokens           []Token
	unresolvedTokens []unresolvedToken
	textStart        int
}

type unresolvedToken struct {
	text              rune
	start             bool
	tokenType         string
	contentTokenStart int
}
type Token struct {
	Text        string     `json:"text"`
	TokenType   string     `json:"tokenType"`
	NodeTagName string     `json:"tagName"`
	NodeClass   string     `json:"class"`
	NodeAttrs   []NodeAttr `json:"attrs"`
}

type NodeAttr struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

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

func updateTokenHtmlByText(t *Token) {
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
	}
}

func appendNewToken(l *Line, t *Token) {
	updateTokenHtmlByText(t)
	l.Tokens = append(l.Tokens, *t)
}

//type token struct {
//    sign string
//    class string
//}

func LinesToHtml(lines [][]Token) string {
	var builder strings.Builder
	builder.WriteString(`<div class="content">`)
	for i := range lines {
		builder.WriteString(lineToHtml(lines[i]))
	}
	builder.WriteString(`</div>`)
	return builder.String()
}

func lineToHtml(tokens []Token) string {
	var builder strings.Builder
	builder.WriteString(`<div class="line">`)
	for i := range tokens {
		builder.WriteString("<")
		builder.WriteString(tokens[i].NodeTagName)
		builder.WriteString(` class="`)
		builder.WriteString(tokens[i].NodeClass)
		builder.WriteString(`" `)
		builder.WriteString(">")
		builder.WriteString(tokens[i].Text)
		//result += "<"+ l.Tokens[i].NodeTagName +" class="++" >"
		builder.WriteString("</")
		builder.WriteString(tokens[i].NodeTagName)
		builder.WriteString(">")
	}
	builder.WriteString(`</div>`)
	return builder.String()
}

// 根据当前line的Tokens生成html
func (l *Line) ToHtml() string {
	return lineToHtml(l.Tokens)
}

// 行内转换的解析函数
func (l *Line) Parse() {
	l.DeleteTextParse()
}

//l.state = LineState.Start
//l.textStart = -1
//for i := 0; i < len(l.Origin); i++ {
//	ch := l.Origin[i]
//	switch l.state {
//	case LineState.Start:
//		l.unresolvedTokens = []unresolvedToken{}
//		switch ch {
//		case '*':
//
//			// 1. 遇到*之后。需要记录这个*留着判断。
//			ut := unresolvedToken{text: '*', start: true, contentTokenStart: len(l.Tokens)}
//			l.unresolvedTokens = append(l.unresolvedTokens, ut)
//			l.state = LineState.ItalicStart
//
//			// 2. 并且还要把*之前的token的text（若有）。
//			if l.textStart != -1 {
//				appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
//				l.textStart = -1
//				l.unresolvedTokens[len(l.unresolvedTokens)-1].contentTokenStart = len(l.Tokens)
//			}
//
//			continue
//		case '~':
//
//			// 1. 遇到*之后。需要记录这个*留着判断。
//			l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '~', start: true, contentTokenStart: len(l.Tokens)})
//			l.state = LineState.DeletedTextStart
//
//			// 2. 并且还要把*之前的token的text（若有）。
//			if l.textStart != -1 {
//				appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
//				l.textStart = -1
//				l.unresolvedTokens[len(l.unresolvedTokens)-1].contentTokenStart = len(l.Tokens)
//			}
//
//			continue
//		default:
//			// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
//			if l.textStart == -1 {
//				l.textStart = i
//			}
//			continue
//		}
//	case LineState.DeletedTextStart:
//		switch ch {
//		case '~':
//			if l.textStart == -1 {
//				if len(l.unresolvedTokens) == 2 {
//					l.textStart = i
//					l.state = LineState.DeletedTextEnd
//				} else {
//					l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '~', start: true})
//				}
//			}
//			continue
//		default:
//			// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
//			if l.textStart == -1 {
//				if len(l.unresolvedTokens) == 2 {
//					l.textStart = i
//					l.state = LineState.DeletedTextEnd
//				} else {
//					l.textStart = i - 1
//					l.unresolvedTokens = l.unresolvedTokens[:len(l.unresolvedTokens)-1]
//					l.state = LineState.Start
//				}
//			} else {
//				if len(l.unresolvedTokens) == 2 {
//					l.state = LineState.DeletedTextEnd
//				} else {
//					l.unresolvedTokens = l.unresolvedTokens[:len(l.unresolvedTokens)-1]
//					l.state = LineState.Start
//				}
//			}
//			continue
//		}
//	case LineState.DeletedTextEnd:
//		switch ch {
//		case '~':
//			length := len(l.unresolvedTokens)
//			if length > 0 {
//				i = l.confirmDeletedText(i)
//				l.state = LineState.Start
//			} else {
//				l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '~', start: true})
//
//				if l.textStart != -1 {
//					appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
//					l.textStart = -1
//				}
//				l.state = LineState.DeletedTextStart
//			}
//			continue
//		default:
//			// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
//			if l.textStart == -1 {
//				l.textStart = i
//			}
//			continue
//		}
//	case LineState.ItalicStart:
//		switch ch {
//		case '*':
//			if l.textStart == -1 {
//				l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '*', start: true})
//			}
//			continue
//		default:
//			// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
//			if l.textStart == -1 {
//				l.textStart = i
//				l.state = LineState.ItalicEnd
//			}
//			continue
//		}
//	case LineState.ItalicEnd:
//		switch ch {
//		case '*':
//			length := len(l.unresolvedTokens)
//			if length > 0 {
//				i = l.confirmItalicType(i)
//				l.state = LineState.Start
//			} else {
//				l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '*', start: true})
//
//				if l.textStart != -1 {
//					appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
//					l.textStart = -1
//				}
//				l.state = LineState.ItalicStart
//			}
//			continue
//		default:
//			// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
//			if l.textStart == -1 {
//				l.textStart = i
//			}
//			continue
//		}
//	}
//
//}
//if len(l.unresolvedTokens) > 0 {
//	//l.Tokens = append(l.Tokens, "*")
//	if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
//		l.Tokens[len(l.Tokens)-1].Text += joinTokens(l.unresolvedTokens, "")
//		updateTokenHtmlByText(&l.Tokens[len(l.Tokens)-1])
//	} else {
//		appendNewToken(l, &Token{Text: joinTokens(l.unresolvedTokens, ""), TokenType: "text"})
//	}
//}
//if l.textStart != -1 {
//	if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
//		l.Tokens[len(l.Tokens)-1].Text += string(l.Origin[l.textStart:])
//		updateTokenHtmlByText(&l.Tokens[len(l.Tokens)-1])
//	} else {
//		appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:]), TokenType: "text"})
//	}
//}

// 斜体、粗体、粗斜体转换方法
func (l *Line) ItalicTextParse() {
	l.state = LineState.Start
	l.textStart = -1
	for i := 0; i < len(l.Origin); i++ {
		ch := l.Origin[i]
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
					appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
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
						appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
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
	l.ResolveLineToken()
}

// 删除线转换方法，会自行调用ItalicTextParse（斜体、粗体、粗斜体转换方法）
func (l *Line) DeleteTextParse() {
	l.state = LineState.Start
	l.textStart = -1
	for i := 0; i < len(l.Origin); i++ {
		ch := l.Origin[i]
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
					appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
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
						appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
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
	l.ResolveLineToken()
	l.ParseWithOther(func(line *Line) {
		line.ItalicTextParse()
	})
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
		appendNewToken(l, &Token{Text: temp, TokenType: tokenType})
		l.textStart = -1
	} else {
		appendNewToken(l, &Token{Text: "", TokenType: tokenType})
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
		appendNewToken(l, &Token{Text: temp, TokenType: tokenType})
		l.textStart = -1
	} else {
		appendNewToken(l, &Token{Text: "", TokenType: tokenType})
	}
	return i
}

// 遍历完成之后将行进行最后的结尾工作：处理未解决的token、处理未发射为text的字符
func (l *Line) ResolveLineToken() {
	if len(l.unresolvedTokens) > 0 {
		//l.Tokens = append(l.Tokens, "*")
		if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
			l.Tokens[len(l.Tokens)-1].Text += joinTokens(l.unresolvedTokens, "")
			updateTokenHtmlByText(&l.Tokens[len(l.Tokens)-1])
		} else {
			appendNewToken(l, &Token{Text: joinTokens(l.unresolvedTokens, ""), TokenType: "text"})
		}
	}
	if l.textStart != -1 {
		if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
			l.Tokens[len(l.Tokens)-1].Text += string(l.Origin[l.textStart:])
			updateTokenHtmlByText(&l.Tokens[len(l.Tokens)-1])
		} else {
			appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:]), TokenType: "text"})
		}
	}
}

// 将当前行的Tokens进行其他Markdown语法的二次转换、转换的方法为参数。转换后，当前行的Tokens将会更新。
func (l *Line) ParseWithOther(parseFunc func(*Line)) {
	tokens := make([]Token, 0)
	for i := range l.Tokens {
		lineText := Line{Origin: []rune(l.Tokens[i].Text), Tokens: []Token{}}
		parseFunc(&lineText)
		tokens = append(tokens, l.Tokens[i].updateWith(lineText.Tokens...)...)
	}
	l.Tokens = tokens
}

func (t Token) updateWith(tokens ...Token) []Token {
	if t.TokenType == "text" {
		return tokens
	}
	for i := range tokens {
		if tokens[i].NodeClass == "text" {
			tokens[i].NodeClass = t.NodeClass
		} else {
			tokens[i].NodeClass += " " + t.NodeClass
		}

		// TODO 出了合并类名之外的更新操作
	}
	return tokens
}
