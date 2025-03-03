package api

import (
	"pacyuribot/api/v1/admin"
	"pacyuribot/api/v1/test"
	"pacyuribot/api/v1/user"
)

var APIGroupApp = new(APIGroup)

type APIGroup struct {
	AdminAPIGroup admin.APIGroup
	TestAPIGroup  test.APIGroup
	UserAPIGroup  user.APIGroup
}
