package router

import (
	"pacyuribot/router/admin"
	"pacyuribot/router/test"
	"pacyuribot/router/user"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Test  test.RouterGroup
	Admin admin.RouterGroup
	User  user.RouterGroup
}
