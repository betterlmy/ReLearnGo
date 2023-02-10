package main

func lengthOfLongestSubstring(s string) int {
	// 检测两个相同字母之间的最长距离
	// abca -> 3
	// asdddasdfgha -> 6
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 1
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			// 检测到重复之后,即发现该子串已达最长,计算此次重复最长的子串距离
			start = lastI + 1 // 从下一个开始进行新的子串检测
		}
		if i-start+1 > maxLength {
			//判断是否超过最长的子串长度
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i

	}
	return maxLength
}
