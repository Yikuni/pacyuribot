package test

import (
	"github.com/gin-gonic/gin"
	"github.com/pluja/pocketbase"
	"os"
	"pacyuribot/global"
	"pacyuribot/logger"
	"pacyuribot/model/common/response"
	"pacyuribot/utils"
)

type PocketBaseAPI struct {
}

type UserLoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	AuthProviders    []interface{} `json:"authProviders"`
	UsernamePassword bool          `json:"usernamePassword"`
	EmailPassword    bool          `json:"emailPassword"`
	OnlyVerified     bool          `json:"onlyVerified"`
	ID               string        `json:"id"`
}

type Data struct {
	Value string `json:"value"`
	Owner string `json:"owner"`
}

type TokenData struct {
	UserID string `json:"userID"`
	Token  string `json:"token"`
}

func (a *PocketBaseAPI) GetToken(c *gin.Context) {
	var info UserLoginInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.InvalidRequestFormat(c)
		logger.Error(err.Error())
		return
	}

	client := pocketbase.NewClient(global.Config.Pocketbase.Url)
	resp, err := pocketbase.CollectionSet[User](client, "users").AuthWithPassword(info.Email, info.Password)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		logger.Error(err.Error())
		return
	}
	logger.Debug("token: %s", resp.Token)
	response.OkWithData(TokenData{
		UserID: resp.Record.ID,
		Token:  resp.Token,
	}, c)
}

func (a *PocketBaseAPI) AuthAndCreateRecord(c *gin.Context) {
	var data TokenData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.InvalidRequestFormat(c)
		logger.Error(err.Error())
		return
	}
	client := pocketbase.NewClient(global.Config.Pocketbase.Url, pocketbase.WithUserToken(data.Token))
	file, err := os.Open("example-config.toml")
	_, err = client.Create("test", map[string]any{
		"value": "test123", "owner": data.UserID, "file": file,
	})
	if err != nil {
		response.FailWithMessage("Failed to create record", c)
		logger.Error("Failed to create record: %s", err.Error())
		return
	}

	response.Ok(c)
}

func (a *PocketBaseAPI) TestAuth(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	id, err := utils.Auth(auth)
	if err != nil {
		response.NoAuth(c)
		return
	}
	response.OkWithMessage(id, c)
}
