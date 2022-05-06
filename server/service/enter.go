package service

import "gin/service/system"

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
