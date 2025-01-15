package crawler

import (
	"pacyuribot/utils"
	"strings"
	"unicode/utf8"
)

type ContentFilter func(content string, d *DefaultCrawler) (string, bool)

func GetTitleFilter(maxLengthC int, maxLengthE int) ContentFilter {
	return func(content string, d *DefaultCrawler) (string, bool) {
		length := utf8.RuneCountInString(content)
		if length <= maxLengthC {
			return "过滤标题", true
		} else if length <= maxLengthE {
			_, english := utils.CheckStringContent(content)
			if english {
				return "过滤英文标题", true
			}
		}
		return content, false
	}
}

func TrimFilter(content string, d *DefaultCrawler) (string, bool) {
	return strings.TrimSpace(content), false
}
