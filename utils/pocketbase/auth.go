package utils

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"io"
	"net/http"
	"pacyuribot/global"
)

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
