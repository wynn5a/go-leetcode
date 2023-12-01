package solution125

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var isAlphaNumeric = func(ch rune) bool {
		return unicode.IsLetter(ch) || unicode.IsNumber(ch)
	}

	// 初始化左右指针
	left, right := 0, len(s)-1
	s = strings.ToLower(s)

	for left < right {
		// 如果左指针所指向的字符不是字母或数字，则左指针右移一位
		if !isAlphaNumeric(rune(s[left])) {
			left++
		} else if !isAlphaNumeric(rune(s[right])) {
			// 如果右指针所指向的字符不是字母或数字，则右指针左移一位
			right--
		} else if s[left] != s[right] {
			// 如果左指针和右指针所指向的字符不相等，则返回false
			return false
		} else {
			// 如果左指针和右指针所指向的字符相等，则同时将左指针右移一位，右指针左移一位
			left++
			right--
		}
	}

	// 当左右指针重合时，表示字符串是回文，返回true
	return true
}
