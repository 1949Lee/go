package parser

import (
	"bytes"
)

// 行内状态类型
type LineStateType int

// 行内状态枚举
type LineStateEnum struct {
	// 开始态
	Start LineStateType

	// 状态1
	State1 LineStateType

	// 状态2
	State2 LineStateType

	// 状态11
	State11 LineStateType

	// 状态12
	State12 LineStateType
}

// 行内状态含义列表
var LineState = LineStateEnum{
	Start:   1,
	State1:  2,
	State2:  3,
	State11: 11,
	State12: 12,
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
	text  rune
	start bool
}
type Token struct {
	Text      string `json:"text"`
	TokenType string `json:"tokenType"`
	Html      string `json:"html"`
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

func appendNewToken(l *Line, t *Token) {
	switch t.TokenType {
	case "text":
		t.Html = "<span class=\"text\">" + t.Text + "</span>"
	case "italic":
		t.Html = "<span class=\"italic\">" + t.Text + "</span>"
	case "bold":
		t.Html = "<span class=\"bold\">" + t.Text + "</span>"
	case "bold-italic":
		t.Html = "<span class=\"bold\">" + t.Text + "</span>"
	}
	l.Tokens = append(l.Tokens, *t)
}

//type token struct {
//    sign string
//    class string
//}

// 行内转换的解析函数
func (l *Line) Parse() {
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
				l.unresolvedTokens = append(l.unresolvedTokens, struct {
					text  rune
					start bool
				}{text: '*', start: true})
				l.state = LineState.State1

				// 2. 并且还要把*之前的token的text（若有）。
				if l.textStart != -1 {
					appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
					l.textStart = -1
				}

				continue
			default:
				// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
					//l.state = LineState.State11
				}
				continue
			}
		case LineState.State1:
			switch ch {
			case '*':
				if l.textStart == -1 {
					l.unresolvedTokens = append(l.unresolvedTokens, struct {
						text  rune
						start bool
					}{text: '*', start: true})
				}
				continue
			default:
				// 如果读到非语法字符，则判断是否为第一个，若是第一个，则记录位置。否则继续扫描下一个字符。
				if l.textStart == -1 {
					l.textStart = i
					l.state = LineState.State11
				}
				continue
			}
		case LineState.State11:
			switch ch {
			case '*':
				l.state = LineState.Start
				length := len(l.unresolvedTokens)
				if length > 0 {
					i = l.confirmTextType(i)
					l.state = LineState.Start
				} else {
					l.unresolvedTokens = append(l.unresolvedTokens, struct {
						text  rune
						start bool
					}{text: '*', start: true})

					if l.textStart != -1 {
						appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
						l.textStart = -1
					}
					l.state = LineState.State1
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
	if len(l.unresolvedTokens) > 0 {
		//l.Tokens = append(l.Tokens, "*")
		if len(l.Tokens) > 0 {
			l.Tokens[len(l.Tokens)-1].Text += joinTokens(l.unresolvedTokens, "")
		} else {
			appendNewToken(l, &Token{Text: joinTokens(l.unresolvedTokens, ""), TokenType: "text"})
		}
	}
	if l.textStart != -1 {
		if len(l.Tokens) > 0 {
			if l.Tokens[len(l.Tokens)-1].TokenType == "text" {
				l.Tokens[len(l.Tokens)-1].Text += string(l.Origin[l.textStart:])
			} else {
				appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:]), TokenType: "text"})
			}
		} else {
			appendNewToken(l, &Token{Text: string(l.Origin[l.textStart:]), TokenType: "text"})
		}
	}
}

// 行内判断token的类型的结束函数，token具体是斜体、加粗、粗斜体
func (l *Line) confirmTextType(i int) int {
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
