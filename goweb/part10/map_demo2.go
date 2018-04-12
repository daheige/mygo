package main

import (
	"fmt"
)

//寻找最长不含有重复字符的子串
func lenOfNonRepeatSubStr(s string) int {
	lastOccured := make(map[rune]int) //一个字符转换为4个字节
	start := 0
	maxLen := 0
	for i, ch := range []rune(s) {
		if last_i, ok := lastOccured[ch]; ok && last_i >= start {
			start = last_i + 1
		}

		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccured[ch] = i
	}

	return maxLen
}

func main() {
	fmt.Println(lenOfNonRepeatSubStr("abcabcbb"))
	fmt.Println(lenOfNonRepeatSubStr("bbb"))
	fmt.Println(lenOfNonRepeatSubStr("一二三"))
	fmt.Println(lenOfNonRepeatSubStr("一二三是你一二"))
}
