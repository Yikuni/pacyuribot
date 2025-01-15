package api

import (
	"pacyuribot/api/v1/admin"
	"pacyuribot/api/v1/test"
)

var APIGroupApp = new(APIGroup)

type APIGroup struct {
	AdminAPIGroup admin.APIGroup
	TestAPIGroup  test.APIGroup
}
