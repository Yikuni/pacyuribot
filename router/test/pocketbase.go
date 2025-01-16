package test

import "github.com/gin-gonic/gin"

type PocketbaseRouter struct {
}

func (s *PocketbaseRouter) InitPocketbaseRouter(rGroup *gin.RouterGroup) {
	r := rGroup.Group("pocketbase")
	r.POST("getToken", pocketbaseAPI.GetToken)
	r.POST("authAndCreateRecord", pocketbaseAPI.AuthAndCreateRecord)
	r.POST("testAuth", pocketbaseAPI.TestAuth)
}
