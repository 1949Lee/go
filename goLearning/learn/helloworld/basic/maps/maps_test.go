package maps

import (
	"testing"
)

// 普通测试
func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{s: "ab这是一个大好人", ans: 9},
		{s: "abcabcabcd", ans: 4},
	}
	for _, tt := range tests {
		if actual := lengthOfLongestSubstring(tt.s); actual != tt.ans {
			t.Errorf("input string %s got %d; expected %d.", tt.s, actual, tt.ans)
		}
	}
}


//性能测试，结果注意单位，ns表示纳秒
func BenchmarkLengthOfLongestSubstring(t *testing.B) {
	s := "ab这是一个大好人"
	ans := 9
	for i :=0;i < 13 ;i++  {
		s = s + s
	}
	t.ResetTimer()

	//性能测试循环测试的次数，由go test自动计算（t.N）
	for i := 0; i < t.N; i++ {
		if actual := lengthOfLongestSubstring(s); actual != ans {
			t.Errorf("input string %s got %d; expected %d.", s, actual, ans)
		}
	}
}
