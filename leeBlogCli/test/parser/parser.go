package parser

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
	Tokens           []string
	unresolvedTokens []unresolvedToken
	textStart        int
}

type unresolvedToken struct {
	text  rune
	start bool
}

//type token struct {
//    sign string
//    class string
//}

func (l *Line) Parse() {
	l.state = LineState.Start
	l.textStart = -1
	for i, ch := range l.Origin {
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
					l.Tokens = append(l.Tokens, string(l.Origin[l.textStart:i]))
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
				if l.textStart != -1 {
					l.Tokens = append(l.Tokens, string(l.Origin[l.textStart:i]))
					l.textStart = -1
				} else {
					l.Tokens = append(l.Tokens, "")
				}
				l.unresolvedTokens = []unresolvedToken{}
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
		l.Tokens = append(l.Tokens, "*")
	}
	if l.textStart != -1 {
		l.Tokens = append(l.Tokens, string(l.Origin[l.textStart:]))
	}
}
