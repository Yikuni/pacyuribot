package utils

import (
	"unicode"
)

// 判断是否是中文字符
func isChineseChar(r rune) bool {
	return unicode.Is(unicode.Han, r)
}

// 判断是否是英文字符
func isEnglishChar(r rune) bool {
	return unicode.IsLetter(r) && r <= unicode.MaxASCII
}

// CheckStringContent 检查字符串全是中文、英文
func CheckStringContent(s string) (allChinese, allEnglish bool) {
	allChinese = true
	allEnglish = true
	for _, r := range s {
		if isChineseChar(r) {
			allEnglish = false
		} else if isEnglishChar(r) {
			allChinese = false
		}
		// 如果两者都已经找到，可以提前退出
		if !allChinese && !allEnglish {
			return
		}
	}
	return
}
