package maps

import "fmt"

// leetcode真题：无重复字符的最长子串 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	start := 0
	currMap := map[rune]int{}
	for i, ch := range []rune(s) {
		if r, ok := currMap[ch]; ok == true && r >= start {
			start = r + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		currMap[ch] = i
	}
	return maxLength
}

func main() {

	// 创建方式一
	m1 := map[string]string{
		"name": "lijiaxuan",
		"age":  "26",
	}
	fmt.Println(m1)

	// 创建方式二
	m2 := make(map[string]string)
	fmt.Println(m2)

	// 空map
	var m3 map[string]int
	fmt.Println(m3)

	// 遍历map,每次遍历输出的顺序可能会不同，因为键值对在map中是无序的，go中的map是哈希map
	for k, v := range m1 {
		//for k, _ := range m1 { // 省略key
		fmt.Println(k, v)
	}

	name, has := m1["name"]
	fmt.Println(name, has)

	// 如果想获取map中不存在的键值对，has会为true。tail会为键值对中值类型的初始值（string为""，int为0）
	tail, has := m1["tail"]
	fmt.Println(tail == "", has)

	//删除map中的键值对,删除不存在的键值对也可以
	delete(m1, "age")
	fmt.Println(m1)

	fmt.Println(lengthOfLongestSubstring("ab这是一个大好人"))
}
