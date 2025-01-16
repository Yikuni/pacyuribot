package router

import (
	"pacyuribot/router/admin"
	"pacyuribot/router/test"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Test  test.RouterGroup
	Admin admin.RouterGroup
}
