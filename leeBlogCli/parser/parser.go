/*
@author 李佳轩
本文件为parser转换markdown语法相关的一些函数和方法。
*/

package parser

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

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
		if ch == '\\' && i < len(l.Origin)-1 && l.Origin[i+1] == markdownRunes.italicRune {
			l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
		}
		switch l.state {
		case LineState.Start:
			l.unresolvedTokens = unresolvedTokenSlice{}
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
		if ch == '\\' && i < len(l.Origin)-1 && l.Origin[i+1] == markdownRunes.deleteRune {
			l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
		}
		switch l.state {
		case LineState.Start:
			l.unresolvedTokens = unresolvedTokenSlice{}
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
	l.unresolvedTokens = unresolvedTokenSlice{}
	for i := 0; i < len(l.Origin); i++ {
		ch := l.Origin[i]
		if ch == '\\' && i < len(l.Origin)-1 && l.Origin[i+1] == markdownRunes.linkRune {
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
						l.unresolvedTokens = unresolvedTokenSlice{}
						l.textStart = -1
					} else {
						// 如果已经读到了行的最后一个字符，则进行一些未完成的token处理
						if i+1 == len(l.Origin) {
							if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
								l.Tokens[len(l.Tokens)-1].Text += "[" + string(l.Origin[l.textStart:]) + "]"
								updateToken(&l.Tokens[len(l.Tokens)-1])
							} else {
								text := ""
								if l.textStart != -1 && l.textStart < len(l.Origin) {
									text = string(l.Origin[l.textStart:])
								}
								l.appendNewToken(&Token{Text: "[" + text + "]", TokenType: "text"})
							}
							l.textStart = -1
							l.unresolvedTokens = unresolvedTokenSlice{}
						}
						l.state = LineState.Start
						l.textStart = -1
						tempToken = Token{}
						l.unresolvedTokens = unresolvedTokenSlice{}
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
					l.unresolvedTokens = unresolvedTokenSlice{}
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
				l.unresolvedTokens = unresolvedTokenSlice{}
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
					l.unresolvedTokens = unresolvedTokenSlice{}
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
		if ch == '\\' && i < len(l.Origin)-1 && l.Origin[i+1] == markdownRunes.backgroundRune {
			l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
		}
		switch l.state {
		case LineState.Start:
			l.unresolvedTokens = unresolvedTokenSlice{}
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
					l.unresolvedTokens = unresolvedTokenSlice{}
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

// 链接转换方法，会自行调用LinkTextParse
func (l *Line) ImageParse() {
	l.state = LineState.Start
	l.textStart = -1
	tempToken := Token{}
	l.unresolvedTokens = unresolvedTokenSlice{}
	for i := 0; i < len(l.Origin); i++ {
		ch := l.Origin[i]
		if ch == '\\' && i < len(l.Origin)-1 && l.Origin[i+1] == markdownRunes.imageRune {
			l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
		}
		switch l.state {
		case LineState.Start:
			switch {
			case ch == '!' && i < len(l.Origin)-1 && l.Origin[i+1] == '[':
				// 1. 遇到[之后。需要记录这个[留着判断。
				l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '!', start: true})
				l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '[', start: true})
				l.state = LineState.ImageTextEnd

				// 2. 并且还要把[之前的token的text（若有）。
				if l.textStart != -1 {
					l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:i]), TokenType: "text"})
					l.textStart = -1
				}
				i++
				continue
			default:
				// 如果读到非语法字符,则记录位置。否则继续扫描下一个字符。textStart == -1表示只记录最开始的位置
				if l.textStart == -1 {
					l.textStart = i
				}
				continue
			}
		case LineState.ImageTextEnd:
			switch ch {
			case ']':
				if len(l.unresolvedTokens) >= 1 {
					// 如果]中有文案，且后面就是(则认为已经进入链接URL判断的状态
					if l.textStart != -1 && i+1 < len(l.Origin) && l.Origin[i+1] == '(' {
						tempToken.Text = string(l.Origin[l.textStart:i])
						tempToken.TokenType = "image"
						l.unresolvedTokens = unresolvedTokenSlice{}
						l.unresolvedTokens = append(l.unresolvedTokens, unresolvedToken{text: '(', start: true})
						l.state = LineState.ImageHrefEnd
						i++
						l.textStart = -1
					} else if l.textStart != -1 && i+1 < len(l.Origin) && l.Origin[i+1] != '(' {
						l.appendNewToken(&Token{Text: l.unresolvedTokens.joinTokens("") + string(l.Origin[l.textStart:i]) + "]", TokenType: "text"})
						l.state = LineState.Start
						l.textStart = -1
						tempToken = Token{}
						l.unresolvedTokens = unresolvedTokenSlice{}
						l.textStart = -1
					} else {
						// 如果已经读到了行的最后一个字符，则进行一些未完成的token处理
						if i+1 == len(l.Origin) {
							if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].TokenType == "text" {
								l.Tokens[len(l.Tokens)-1].Text += "[" + string(l.Origin[l.textStart:]) + "]"
								updateToken(&l.Tokens[len(l.Tokens)-1])
							} else {
								text := ""
								if l.textStart != -1 && l.textStart < len(l.Origin) {
									text = string(l.Origin[l.textStart:])
								}
								l.appendNewToken(&Token{Text: "[" + text + "]", TokenType: "text"})
							}
							l.textStart = -1
							l.unresolvedTokens = unresolvedTokenSlice{}
						}
						l.state = LineState.Start
						l.textStart = -1
						tempToken = Token{}
						l.unresolvedTokens = unresolvedTokenSlice{}
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

						l.Tokens[len(l.Tokens)-1].Text += l.unresolvedTokens.joinTokens("") + string(l.Origin[l.textStart:])
						updateToken(&l.Tokens[len(l.Tokens)-1])
					} else {
						l.appendNewToken(&Token{Text: l.unresolvedTokens.joinTokens("") + string(l.Origin[l.textStart:]), TokenType: "text"})
					}
					l.textStart = -1
					l.unresolvedTokens = unresolvedTokenSlice{}
				}
				continue
			}
		case LineState.ImageHrefEnd:
			switch ch {
			case ')':
				length := len(l.unresolvedTokens)
				if length == 1 {
					if l.textStart != -1 {
						tempToken.NodeAttrs = []NodeAttr{
							{Key: "src", Value: string(l.Origin[l.textStart:i])},
						}
						l.appendNewToken(&tempToken)
					} else {
						var builder strings.Builder
						builder.WriteString("![")
						builder.WriteString(tempToken.Text)
						builder.WriteString("]()")
						l.appendNewToken(&Token{Text: builder.String(), TokenType: "text"})
					}
				}
				l.state = LineState.Start
				l.textStart = -1
				tempToken = Token{}
				l.unresolvedTokens = unresolvedTokenSlice{}
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
						text := ""
						if l.textStart != -1 && l.textStart < len(l.Origin) {
							text = string(l.Origin[l.textStart:])
						}
						l.appendNewToken(&Token{Text: "[" + tempToken.Text + "](" + text, TokenType: "text"})
					}
					l.textStart = -1
					l.unresolvedTokens = unresolvedTokenSlice{}
				}
				continue
			}
		}

	}
	l.resolveLineToken()
	l.parseWithOther(func(line *Line) {
		line.BackgroundStrongParse()
	})
}

//多级标题转换方法，会自行调用ImageParse
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
						l.ImageParse()
						return
					}
				}
				if ch == '\\' && i < len(l.Origin)-1 && l.Origin[i+1] == markdownRunes.headerRune {
					l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
				}
			}
		}
		if l.textStart < len(l.Origin) && l.textStart != -1 {
			l.appendNewToken(&Token{Text: string(l.Origin[l.textStart:]), NodeTagName: "h" + strconv.Itoa(headerLevel), TokenType: "header"})
			l.parseWithOther(func(line *Line) {
				line.ImageParse()
			})
		}
	} else { // 行开头不是#，直接进行其他转换
		for i := 0; i < len(l.Origin); i++ {
			ch := l.Origin[i]
			if ch == '\\' && i < len(l.Origin)-1 && l.Origin[i+1] == markdownRunes.headerRune {
				l.Origin = append(l.Origin[:i], l.Origin[i+1:]...)
			}
		}
		l.ImageParse()
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
			temp = l.unresolvedTokens.joinTokens("") + string(l.Origin[l.textStart:originIndex])

			//制空开始数组。
			l.unresolvedTokens = unresolvedTokenSlice{}
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
			l.Tokens[len(l.Tokens)-1].Text += l.unresolvedTokens.joinTokens("")
			updateToken(&l.Tokens[len(l.Tokens)-1])
		} else {
			l.appendNewToken(&Token{Text: l.unresolvedTokens.joinTokens(""), TokenType: "text"})
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
	tokens := make(TokenSlice, 0)
	for i := range l.Tokens {
		lineText := Line{Origin: []rune(l.Tokens[i].Text), Tokens: TokenSlice{}}
		parseFunc(&lineText)
		tokens = append(tokens, l.Tokens[i].updateWith(lineText.Tokens...)...)
	}
	l.Tokens = tokens
}

// 用新的token去更新某个token
func (t Token) updateWith(tokens ...Token) TokenSlice {
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
		return TokenSlice{t}
	case "image":
		t.Children = tokens
		t.Text = ""
		return TokenSlice{t}
	}
	return tokens
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
	case "image":
		t.NodeClass = "image"
		t.NodeTagName = "img"
		t.NodeAttrs = append(t.NodeAttrs, NodeAttr{Key: "alt", Value: ""})
	case "header":
		t.NodeClass = "header-" + t.NodeTagName
	}
}

// 将传入的行转换为html
func LinesToHtml(lines []TokenSlice) string {
	var builder strings.Builder
	builder.WriteString(`<div class="content">`)
	for i := range lines {
		builder.WriteString(lineToHtml(lines[i]))
	}
	builder.WriteString(`</div>`)
	return builder.String()
}

// 将传入的token数组转化为html
func lineToHtml(tokens TokenSlice) string {
	var builder strings.Builder
	s := TokenSlice(tokens)
	if tokens[0].TokenType == "empty-line-br" {
		builder.WriteString(`<div class="block empty-single-line-block">`)
		builder.WriteString(tokensToHtml(tokens))
		builder.WriteString(`</div>`)
	} else if tokens[0].TokenType == "styled-block" {
		tokens[0].NodeClass += " block"
		builder.WriteString(tokensToHtml(tokens))
	} else if ok, index := s.has("image"); ok {
		if len(s) == 1 {
			builder.WriteString(`<div class="block image-block">`)
			builder.WriteString(tokensToHtml(tokens))
			builder.WriteString(`</div>`)
		} else {
			// image之前
			if index > 0 {
				builder.WriteString(`<div class="block">`)
				builder.WriteString(tokensToHtml(tokens[:index]))
				builder.WriteString(`</div>`)
			}
			builder.WriteString(`<div class="block image-list-block">`)
			builder.WriteString(tokensToHtml(TokenSlice{tokens[index]}))
			builder.WriteString(`</div>`)
			if index < len(s)-1 {
				builder.WriteString(`<div class="block">`)
				builder.WriteString(tokensToHtml(tokens[index+1:]))
				builder.WriteString(`</div>`)
			}
		}
	} else {
		builder.WriteString(`<div class="block">`)
		builder.WriteString(tokensToHtml(tokens))
		builder.WriteString(`</div>`)
	}
	return builder.String()
}

// 将token转化为html
func tokensToHtml(tokens TokenSlice) string {
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
		} else if tokens[i].TokenType == "image" { // image的特殊处理
			builder.WriteString(`<div class="image-wrapper">`)
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
			builder.WriteString(`<div class="image-text-wrapper">`)
			builder.WriteString(tokens[i].Text)
			if len(tokens[i].Children) > 0 {
				builder.WriteString(tokensToHtml(tokens[i].Children))
			}
			builder.WriteString(`</div></div>`)
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

// 检查开头的缩进的空格数目，并将去掉空格后的真是内容返回
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
	if lineText == "" || realRune == "" {
		return false, BlockResult{}
	}
	switch realRune[0] {
	case '*', '-': // 无序列表list和check-list
		if len(realRune) >= 2 && realRune[1] == ' ' {
			if len(realRune) >= 5 && realRune[2] == '[' {
				if (realRune[3] == 'x' || realRune[3] == 'X') && realRune[4] == ']' { // 勾选的情况
					return true, BlockResult{TokenType: "check-list", IndentCount: indentCount, CheckBoxStatus: true}
				} else if realRune[3] == ']' || (realRune[3] == ' ' && realRune[4] == ']') { // 未勾选的情况
					return true, BlockResult{TokenType: "check-list", IndentCount: indentCount, CheckBoxStatus: false}
				} else {
					return true, BlockResult{TokenType: "list", IndentCount: indentCount}
				}
			} else {
				return true, BlockResult{TokenType: "list", IndentCount: indentCount}
			}
		}
	case '>': // 样式块styled-block或文字引用block-quote
		if len(realRune) >= 3 && realRune[1] == '>' && realRune[2] == '>' {
			return true, BlockResult{TokenType: "styled-block"}
		} else {
			return true, BlockResult{TokenType: "block-quote"}
		}
	case '|': // 表格table
		return true, BlockResult{TokenType: "table-block"}
	case '`': // 代码块code-block
		if len(realRune) > 2 && realRune[1] == '`' && realRune[2] == '`' {
			return true, BlockResult{TokenType: "code-block"}
		}
	case '+': // 自动有序列表auto-order-list
		if len(realRune) > 1 && realRune[1] == ' ' {
			return true, BlockResult{TokenType: "auto-order-list", IndentCount: indentCount}
		}
	case ':': // 名词定义列表word-list
		if len(realRune) > 1 && realRune[1] == ':' {
			return true, BlockResult{TokenType: "word-list"}
		}
	}
	return false, BlockResult{TokenType: "text", IndentCount: indentCount}
}

func markdownLinesParse(str string, blockResult BlockResult, dataList *[]TokenSlice) {
	line := Line{
		Origin: []rune(str),
		Tokens: TokenSlice{}}
	var tokens TokenSlice
	if blockResult.IndentCount == 4 {
		line.Origin = []rune(str[blockResult.IndentCount:])
		line.LineParse()
		tokens = TokenSlice{
			{
				TokenType:   "text-indent-tag",
				NodeTagName: "span",
				NodeClass:   "text-indent-tag",
				Text:        "缩进",
			},
		}
		tokens = append(tokens, line.Tokens...)
	} else {
		line.LineParse()
		if len(line.Tokens) > 0 {
			tokens = line.Tokens
		}
	}
	if tokens != nil {
		*dataList = append(*dataList, tokens)
	}
}

// 接受markdown字符串，并将之转化为html
func MarkdownParse(markdownText string) ([]TokenSlice, string) {
	//scanner := bufio.NewScanner(strings.NewReader(markdownText))
	//dataList := make([]TokenSlice, 0)
	//for scanner.Scan() {
	//	lineText := scanner.Text()
	//	line := Line{Origin: []rune(lineText), Tokens: TokenSlice{}}
	//	line.LineParse()
	//	dataList = append(dataList, line.Tokens)
	//}

	//这种split的方法比bufio那种读取块100-500微秒。
	list := strings.Split(markdownText, "\n")
	dataList := make([]TokenSlice, 0)
	for i := 0; i < len(list); i++ {
		if ok, blockResult := isInBlock(list[i]); ok {
			var tokens TokenSlice
			i, tokens = blockParse(list, i, blockResult)
			if tokens == nil {
				markdownLinesParse(list[i], blockResult, &dataList)
			} else {
				if i < len(list) {
					i--
				}
				dataList = append(dataList, tokens)
			}
		} else {
			markdownLinesParse(list[i], blockResult, &dataList)
		}
	}
	return dataList, LinesToHtml(dataList)
}

// 将多行转换为块。
func blockParse(lines []string, index int, blockResult BlockResult) (int, TokenSlice) {
	var (
		tokens TokenSlice
		i      int
	)
	switch blockResult.TokenType {
	case "list":
		i, tokens = listParse(lines, index, blockResult)
		return i, tokens
	case "check-list":
		i, tokens = checkListParse(lines, index, blockResult)
		return i, tokens
	case "auto-order-list":
		i, tokens = autoOrderListParse(lines, index, blockResult, nil)
		return i, tokens
	case "block-quote":
		i, tokens = blockQuoteParse(lines, index, blockResult)
		return i, tokens
	case "styled-block":
		i, tokens = styledBlockParse(lines, index, blockResult)
		return i, tokens
	case "code-block":
		i, tokens = codeBlockParse(lines, index, blockResult)
		return i, tokens
	case "table-block":
		i, tokens = tableBlockParse(lines, index, blockResult)
		return i, tokens
	}
	return index, nil
}

// 有序列表转换方法
func autoOrderListParse(lines []string, index int, blockResult BlockResult, level []string) (int, TokenSlice) {
	originLevel := blockResult.IndentCount/4 + 1
	tokens := TokenSlice{
		{
			TokenType:   "auto-order-list",
			NodeTagName: "ol",
			NodeClass:   "auto-order-list list-level-" + strconv.Itoa(originLevel),
			Children:    TokenSlice{},
		}}
	originIndent := blockResult.IndentCount

	i := index
	for ; i < len(lines); i++ {
		ok, temResult := isInBlock(lines[i])
		if ok && temResult.TokenType == "auto-order-list" { // list的新一行
			if temResult.IndentCount >= originIndent && temResult.IndentCount < originIndent+4 { // list的新一行
				tokens[0].Children = append(tokens[0].Children, Token{
					TokenType:   "auto-order-list-item",
					NodeTagName: "li",
					NodeClass:   "auto-order-list-item list-item-level-" + strconv.Itoa(originLevel),
					Children: TokenSlice{
						{TokenType: "auto-order-list-item-text-line-wrapper", NodeTagName: "div", NodeClass: "auto-order-list-item-text-line-wrapper", Children: TokenSlice{}},
					},
				})
				text := []rune(lines[i])[2+temResult.IndentCount:]
				line := Line{Origin: text, Tokens: TokenSlice{}}
				line.LineParse()
				indexToken := Token{
					TokenType:   "auto-order-list-item-level-index",
					NodeTagName: "span",
					NodeClass:   "auto-order-list-item-level-index",
				}
				if level != nil {
					tempLevel := append(level, strconv.Itoa(len(tokens[0].Children)))
					indexToken.Text = strings.Join(tempLevel, AutoOrderListLevelIndexDivider)
				} else {
					indexToken.Text = strconv.Itoa(len(tokens[0].Children)) + AutoOrderListLevelIndexDivider
				}
				tokens[0].Children[len(tokens[0].Children)-1].Children[0].Children = append(TokenSlice{indexToken}, line.Tokens...)
			} else if temResult.IndentCount >= originIndent+4 { //新的列表项，缩进符合下一级。
				//var temIndex int
				var subTokens TokenSlice
				if level != nil {
					i, subTokens = autoOrderListParse(lines, i, temResult, append(level, strconv.Itoa(len(tokens[0].Children))))
				} else {
					i, subTokens = autoOrderListParse(lines, i, temResult, []string{strconv.Itoa(len(tokens[0].Children))})
				}
				i--
				tokens[0].Children[len(tokens[0].Children)-1].Children = append(tokens[0].Children[len(tokens[0].Children)-1].Children, subTokens...)
			} else if temResult.IndentCount < originIndent {
				return i, tokens
			}
		} else if !ok && temResult.TokenType == "text" {
			ci := len(tokens[0].Children) - 1
			tokens[0].Children[ci].Children = append(tokens[0].Children[ci].Children, Token{
				TokenType:   "auto-order-list-item-text-line-wrapper",
				NodeTagName: "div",
				NodeClass:   "auto-order-list-item-text-line-wrapper",
				Children:    TokenSlice{}})
			di := len(tokens[0].Children[ci].Children) - 1
			text := []rune(strings.Replace(lines[i][originIndent:], " ", "\u2002", temResult.IndentCount))
			line := Line{Origin: text, Tokens: TokenSlice{}}
			line.LineParse()
			tokens[0].Children[ci].Children[di].Children = line.Tokens
		} else if !ok { // list结束
			index = i
			return index, tokens
		} else { // 其他块
			if temResult.IndentCount >= originIndent+4 {
				temTokens := make(TokenSlice, 0)
				i, temTokens = blockParse(lines, i, temResult)
				tokens[0].Children[len(tokens[0].Children)-1].Children = append(tokens[0].Children[len(tokens[0].Children)-1].Children, temTokens...)
				i--
			} else {
				index = i
				return index, tokens
			}
		}
	}
	index = i
	return index, tokens
}

// check列表
func checkListParse(lines []string, index int, blockResult BlockResult) (int, TokenSlice) {
	originLevel := blockResult.IndentCount/4 + 1
	tokens := TokenSlice{
		{
			TokenType:   "check-list",
			NodeTagName: "ul",
			NodeClass:   "check-list check-list-level-" + strconv.Itoa(originLevel),
			Children:    TokenSlice{},
		}}
	originIndent := blockResult.IndentCount

	i := index
	for ; i < len(lines); i++ {
		ok, temResult := isInBlock(lines[i])
		if ok && temResult.TokenType == "check-list" { // check-list的新一行
			if temResult.IndentCount >= originIndent && temResult.IndentCount < originIndent+4 { // check-list的新一行
				tokens[0].Children = append(tokens[0].Children, Token{
					TokenType:   "check-list-item",
					NodeTagName: "li",
					NodeClass:   "check-list-item check-list-item-level-" + strconv.Itoa(originLevel),
					Children: TokenSlice{
						{TokenType: "check-list-item-text-line-wrapper", NodeTagName: "div", NodeClass: "check-list-item-text-line-wrapper title", Children: TokenSlice{}},
					},
				})
				checkBox, checkIndent := getCheckToken(temResult.CheckBoxStatus)
				text := []rune(lines[i])[2+temResult.IndentCount+checkIndent:]
				line := Line{Origin: text, Tokens: TokenSlice{}}
				line.LineParse()
				tokens[0].Children[len(tokens[0].Children)-1].Children[0].Children = append(TokenSlice{checkBox}, line.Tokens...)
			} else if temResult.IndentCount >= originIndent+4 { //新的列表项，缩进符合下一级。
				temIndex, subTokens := checkListParse(lines, i, temResult)
				i = temIndex - 1
				tokens[0].Children[len(tokens[0].Children)-1].Children = append(tokens[0].Children[len(tokens[0].Children)-1].Children, subTokens...)
			} else if temResult.IndentCount < originIndent {
				return i, tokens
			}
		} else if !ok && temResult.TokenType == "text" {
			ci := len(tokens[0].Children) - 1
			tokens[0].Children[ci].Children = append(tokens[0].Children[ci].Children, Token{
				TokenType:   "check-list-item-text-line-wrapper",
				NodeTagName: "div",
				NodeClass:   "check-list-item-text-line-wrapper",
				Children:    TokenSlice{}})
			di := len(tokens[0].Children[ci].Children) - 1
			text := []rune(strings.Replace(lines[i][originIndent:], " ", "\u2002", temResult.IndentCount))
			line := Line{Origin: text, Tokens: TokenSlice{}}
			line.LineParse()
			tokens[0].Children[ci].Children[di].Children = line.Tokens
		} else { // list结束
			index = i
			return index, tokens
		}
	}
	index = i
	return index, tokens
}

// 获取check列表的
func getCheckToken(status bool) (checkBox Token, length int) {
	length = 3
	checkBox = Token{}
	checkBoxAttrs := []NodeAttr{
		{
			Key:   "type",
			Value: "checkbox",
		},
		{
			Key:   "checked",
			Value: "false",
		},
	}
	if status {
		checkBoxAttrs[1].Value = "true"
	}
	checkBox.NodeTagName = "input"
	checkBox.NodeAttrs = checkBoxAttrs
	checkBox.NodeClass = "check-list-checkbox"
	checkBox.TokenType = "check-list-checkbox"
	return checkBox, length
}

// 无序列表转换方法
func listParse(lines []string, index int, blockResult BlockResult) (int, TokenSlice) {
	originLevel := blockResult.IndentCount/4 + 1
	tokens := TokenSlice{
		{
			TokenType:   "list",
			NodeTagName: "ul",
			NodeClass:   "list list-level-" + strconv.Itoa(originLevel),
			Children:    TokenSlice{},
		}}
	originIndent := blockResult.IndentCount

	i := index
	for ; i < len(lines); i++ {
		ok, temResult := isInBlock(lines[i])
		if ok && temResult.TokenType == "list" { // list的新一行
			if temResult.IndentCount >= originIndent && temResult.IndentCount < originIndent+4 { // list的新一行
				tokens[0].Children = append(tokens[0].Children, Token{
					TokenType:   "list-item",
					NodeTagName: "li",
					NodeClass:   "list-item list-item-level-" + strconv.Itoa(originLevel),
					Children: TokenSlice{
						{TokenType: "list-item-text-line-wrapper", NodeTagName: "div", NodeClass: "list-item-text-line-wrapper title", Children: TokenSlice{}},
					},
				})
				text := []rune(lines[i])[2+temResult.IndentCount:]
				line := Line{Origin: text, Tokens: TokenSlice{}}
				line.LineParse()
				tokens[0].Children[len(tokens[0].Children)-1].Children[0].Children = line.Tokens
			} else if temResult.IndentCount >= originIndent+4 { //新的列表项，缩进符合下一级。
				temIndex, subTokens := listParse(lines, i, temResult)
				i = temIndex - 1
				tokens[0].Children[len(tokens[0].Children)-1].Children = append(tokens[0].Children[len(tokens[0].Children)-1].Children, subTokens...)
			} else if temResult.IndentCount < originIndent {
				return i, tokens
			}
		} else if !ok && temResult.TokenType == "text" {
			ci := len(tokens[0].Children) - 1
			tokens[0].Children[ci].Children = append(tokens[0].Children[ci].Children, Token{
				TokenType:   "list-item-text-line-wrapper",
				NodeTagName: "div",
				NodeClass:   "list-item-text-line-wrapper",
				Children:    TokenSlice{}})
			di := len(tokens[0].Children[ci].Children) - 1
			text := []rune(strings.Replace(lines[i][originIndent:], " ", "\u2002", temResult.IndentCount))
			line := Line{Origin: text, Tokens: TokenSlice{}}
			line.LineParse()
			tokens[0].Children[ci].Children[di].Children = line.Tokens
		} else { // list结束
			index = i
			return index, tokens
		}
	}
	index = i
	return index, tokens
}

// 文字引用块转换方法
func blockQuoteParse(lines []string, index int, _ BlockResult) (int, TokenSlice) {
	tokens := TokenSlice{
		{
			TokenType:   "block-quote",
			NodeTagName: "div",
			NodeClass:   "block-quote",
			Children:    TokenSlice{},
		}}
	//originIndent := blockResult.IndentCount

	i := index
	for ; i < len(lines); i++ {
		ok, temResult := isInBlock(lines[i])
		if ok && temResult.TokenType == "block-quote" { // 是新的一行。
			tokens[0].Children = append(tokens[0].Children, Token{
				TokenType:   "block-quote-line",
				NodeTagName: "div",
				NodeClass:   "block-quote-line",
			})
			if len(lines[i]) > 2 {
				if subOk, subResult := isInBlock(lines[i][2:]); subOk {
					temTokens := make(TokenSlice, 0)
					var subI int
					var subLines []string
					l := i
					for ; l < len(lines); l++ {
						if len(lines[l]) >= 2 {
							subLines = append(subLines, lines[l][2:])
						} else {
							break
						}
					}
					subI, temTokens = blockParse(subLines, 0, subResult)
					i = i + subI
					tokens[0].Children[len(tokens[0].Children)-1].Children = append(tokens[0].Children[len(tokens[0].Children)-1].Children, temTokens...)
					i--
				} else {
					if 2+temResult.IndentCount < len([]rune(lines[i])) {
						text := []rune(lines[i])[2+temResult.IndentCount:]
						line := Line{Origin: text, Tokens: TokenSlice{}}
						line.LineParse()
						tokens[0].Children[len(tokens[0].Children)-1].Children = line.Tokens
					}
				}
			}
		} else { // 引用块结束
			index = i
			return index, tokens
		}
	}
	index = i
	return index, tokens
}

// 颜色块转换方法
func styledBlockParse(lines []string, index int, _ BlockResult) (int, TokenSlice) {
	tokens := TokenSlice{
		{
			TokenType:   "styled-block",
			NodeTagName: "div",
			NodeClass:   "styled-block",
			Children:    TokenSlice{},
		}}
	//originIndent := blockResult.IndentCount

	i := index
	re := regexp.MustCompile(`>+\s*(.+)`)
	tokens[0].NodeAttrs = handleTokenNodeInlineStyle(re.FindAllStringSubmatch(lines[i], -1)[0][1])
	i++
	for ; i < len(lines); i++ {
		ok, temResult := isInBlock(lines[i])
		if ok && temResult.TokenType == "styled-block" { // 样式块结束。
			index = i + 1
			return index, tokens
		} else if !ok { // 非块行
			tokens[0].Children = append(tokens[0].Children, Token{
				TokenType:   "styled-block-line",
				NodeTagName: "div",
				NodeClass:   "styled-block-line",
			})
			line := Line{Origin: []rune(lines[i]), Tokens: TokenSlice{}}
			line.LineParse()
			tokens[0].Children[len(tokens[0].Children)-1].Children = line.Tokens
		} else { // 其他块
			tokens[0].Children = append(tokens[0].Children, Token{
				TokenType:   "styled-block-line",
				NodeTagName: "div",
				NodeClass:   "styled-block-line",
			})
			temTokens := make(TokenSlice, 0)
			i, temTokens = blockParse(lines, i, temResult)
			tokens[0].Children[len(tokens[0].Children)-1].Children = append(tokens[0].Children[len(tokens[0].Children)-1].Children, temTokens...)
			i--
		}
	}
	index = i
	return index, tokens
}

// 代码块儿转换
func codeBlockParse(lines []string, index int, _ BlockResult) (int, TokenSlice) {
	tokens := TokenSlice{
		{
			TokenType:   "code-block",
			NodeTagName: "pre",
			NodeClass:   "code-block",
			Children: TokenSlice{
				{
					TokenType:   "code-block-code",
					NodeTagName: "code",
				},
			},
		}}
	var buffer bytes.Buffer
	i := index
	re := regexp.MustCompile("`+\\s*([^`]+)")
	lang := re.FindAllStringSubmatch(lines[i], -1)
	// 判断是否输入语言
	if lang == nil {
		return index, nil
	}
	tokens[0].Children[0].NodeClass = lang[0][1]
	i++
	for ; i < len(lines); i++ {
		ok, temResult := isInBlock(lines[i])
		if ok && temResult.TokenType == "code-block" { // 代码块结束。
			index = i + 1
			break
		} else {
			buffer.WriteString(lines[i])
			buffer.WriteString("\n")
		}
	}
	// 判断是否检测到代码块的结束标志。
	if i == len(lines) {
		// 没有检测到代码块儿结束
		return index, nil
	}
	tokens[0].Children[0].Text = buffer.String()
	return index, tokens
}

// 表格块转换
func tableBlockParse(lines []string, index int, _ BlockResult) (int, TokenSlice) {
	tokens := TokenSlice{
		{
			TokenType:   "table-block",
			NodeTagName: "table",
			NodeClass:   "table-block",
			Children: TokenSlice{
				// 表头
				{
					TokenType:   "table-block-thead",
					NodeTagName: "thead",
					NodeClass:   "table-block-thead",
					Children: TokenSlice{
						// 表头唯一一行
						{
							TokenType:   "table-block-thead-tr",
							NodeTagName: "tr",
							NodeClass:   "table-block-thead-tr",
							Children:    TokenSlice{
								// 表头每一单元格
							},
						},
					},
				},

				// 表体
				{
					TokenType:   "table-block-tbody",
					NodeTagName: "tbody",
					NodeClass:   "table-block-tbody",
					Children:    TokenSlice{},
				},
			},
		}}
	//var buffer bytes.Buffer
	i := index
	temStr := strings.Split(lines[i], "|")
	colRegTemp := temStr[1 : len(temStr)-1]
	i++

	// 表格列的数量
	colCount := len(colRegTemp)
	colTextAlign := getTableColTextAlign(colCount, lines[i])
	i++
	for j := 0; j < colCount; j++ {
		thToken := Token{
			TokenType:   "table-block-thead-th",
			NodeTagName: "th",
			NodeClass:   "table-block-thead-th",
			Text:        strings.TrimRight(colRegTemp[j], " "),
		}
		if colTextAlign[j] != "" {
			thToken.NodeAttrs = []NodeAttr{
				{
					Key:   "style",
					Value: colTextAlign[j],
				},
			}
		}
		// 将表格标题单元格的文案继续转换
		line := Line{Origin: []rune(thToken.Text), Tokens: TokenSlice{}}
		line.LineParse()
		thToken.Children = line.Tokens
		thToken.Text = ""
		tokens[0].Children[0].Children[0].Children = append(tokens[0].Children[0].Children[0].Children, thToken)
	}
	for ; i < len(lines); i++ {
		ok, temResult := isInBlock(lines[i])
		if ok && temResult.TokenType == "table-block" { // 表格新的一行。
			tokens[0].Children[1].Children = append(tokens[0].Children[1].Children, Token{
				TokenType:   "table-block-tbody-tr",
				NodeTagName: "tr",
				NodeClass:   "table-block-tbody-tr",
				Children:    TokenSlice{}})
			currTrIndex := len(tokens[0].Children[1].Children) - 1
			tdTemStr := strings.Split(lines[i], "|")
			tdRegTemp := tdTemStr[1 : len(tdTemStr)-1]
			for j := 0; j < len(tdRegTemp); j++ {
				tdToken := Token{
					TokenType:   "table-block-tbody-td",
					NodeTagName: "td",
					NodeClass:   "table-block-tbody-td",
					Text:        strings.TrimRight(tdRegTemp[j], " "),
				}
				//if strings.
				if attrs, str := getSpan(tdToken.Text); attrs != nil {
					tdToken.NodeAttrs = append(tdToken.NodeAttrs, attrs...)
					tdToken.Text = str
				}
				if colTextAlign[j] != "" {
					tdToken.NodeAttrs = append(tdToken.NodeAttrs, NodeAttr{
						Key:   "style",
						Value: colTextAlign[j],
					})
				}
				// 将表格标题单元格的文案继续转换
				line := Line{Origin: []rune(tdToken.Text), Tokens: TokenSlice{}}
				line.LineParse()
				tdToken.Children = line.Tokens
				tdToken.Text = ""
				tokens[0].Children[1].Children[currTrIndex].Children = append(tokens[0].Children[1].Children[currTrIndex].Children, tdToken)
			}
		} else {
			index = i
			return index, tokens
		}
	}
	//tokens[0].Children[0].Text = buffer.String()
	index = i
	return index, tokens
}

//是否含有colspan或rowspan，有则返回其值
func getSpan(str string) ([]NodeAttr, string) {
	re := regexp.MustCompile(`:-(\d+)-:|:\+(\d)+\+:`)
	temp := re.FindAllStringSubmatch(str, -1)
	str = re.ReplaceAllString(str, "")
	attrs := make([]NodeAttr, 0)
	if temp != nil {
		for _, v := range temp {
			if v[1] != "" {
				attrs = append(attrs, NodeAttr{Key: "colspan", Value: v[1]})
			}

			if v[2] != "" {
				attrs = append(attrs, NodeAttr{Key: "rowspan", Value: v[2]})
			}
		}
		return attrs, str
	} else {
		return nil, str
	}

}

// 根据表格第二行的配置，得到每列的对齐样式。
func getTableColTextAlign(colCount int, line string) []string {

	temStr := strings.Split(line, "|")
	colRegTemp := temStr[1 : len(temStr)-1]
	colTextAlign := make([]string, colCount)
	for i := 0; i < colCount; i++ {

		// 左对齐
		if strings.Contains(colRegTemp[i], "-:") {

			// 居中
			if strings.Contains(colRegTemp[i], ":-") {
				colTextAlign[i] = "text-align:center;"
				continue
			} else {
				// 左对齐
				colTextAlign[i] = "text-align:left;"
				continue
			}
		}

		// 右对齐
		if strings.Contains(colRegTemp[i], ":-") {
			colTextAlign[i] = "text-align:right;"
			continue
		}
	}
	return colTextAlign
}

// 根据传入的字符串，解析并返回得到的html node attr
func handleTokenNodeInlineStyle(str string) []NodeAttr {
	finalAttrs := make([]NodeAttr, 0)
	finalAttrs = append(finalAttrs, NodeAttr{Key: "style", Value: str})
	return finalAttrs
}
