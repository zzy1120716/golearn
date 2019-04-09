// map应用示例
// LeetCode #3
// 寻找最长不含有重复字符的子串（国际版）
// abcabcbb --> abc
// bbbbb --> b
// pwwkew --> wke
package nonrepeating

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		// 如果ch不在map中，lastOccurred[ch]值为0
		// 是合法下标，但会出问题，因此创建变量lastI，先进行判断
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
