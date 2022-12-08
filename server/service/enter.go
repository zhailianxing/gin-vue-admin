package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Test"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	TestServiceGroup    Test.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
