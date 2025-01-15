package utils

import (
	"fmt"
	"net/url"
)

// CompleteURL 补全URL，返回完整的URL
func CompleteURL(inputURL string, baseURL *url.URL) (*url.URL, error) {
	// 如果是 # 开头的锚点，直接返回空字符串
	if len(inputURL) > 0 && inputURL[0] == '#' {
		return nil, nil
	}

	// 解析 inputURL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return nil, fmt.Errorf("invalid input URL: %w", err)
	}

	// 使用 ResolveReference 方法补全URL
	resolvedURL := baseURL.ResolveReference(parsedURL)

	return resolvedURL, nil
}
