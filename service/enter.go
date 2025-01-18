package service

import "pacyuribot/service/admin"

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	AdminServiceGroup admin.AdminServiceGroup
}
