package utils

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
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

func Auth(token string) (string, error) {
	url := global.Config.Pocketbase.Url + "/api/collections/users/auth-refresh"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	res, err := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unauthorized")
	}
	body, _ := io.ReadAll(res.Body)
	jsonObj, _ := gabs.ParseJSON(body)
	id, _ := jsonObj.Path("record.id").Data().(string)
	return id, nil
}
