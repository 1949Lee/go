/*
@author 李佳轩
本文件为parser中用到的切片的相关方法。
*/

package parser

import "bytes"

// 判断token数组中第一项TokenType为特定值的下标。若没找到，返回true和-1。若找到了，返回true和第一项的下标
func (s TokenSlice) has(tokenType string) (bool, int) {
	for i, t := range s {
		if t.TokenType == tokenType {
			return true, i
		}
	}
	return false, -1
}

// 将传入的token数组转化为字符串，数组中的每一项转化为字符串之后用str连接起来形成一个新的字符串
func (p unresolvedTokenSlice) joinTokens(str string) string {
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
