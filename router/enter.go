package router

import "pacyuribot/router/test"

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Test test.RouterGroup
}
