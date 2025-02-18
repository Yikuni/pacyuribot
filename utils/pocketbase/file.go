package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"pacyuribot/global"
)

func GetPocketbaseFile(collectionID string, fileID string, fileName string) (string, error) {
	fileURL := fmt.Sprintf("%s/api/files/%s/%s/%s", global.Config.Pocketbase.Url, collectionID, fileID, fileName)
	resp, err := http.Get(fileURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	err = os.MkdirAll("cache/data_file", os.ModePerm.Perm())
	if err != nil {
		return "", err
	}
	// 创建本地文件
	file, err := os.Create("cache/data_file/" + fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 将HTTP响应的内容写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return "cache/data_file/" + fileName, err
}
