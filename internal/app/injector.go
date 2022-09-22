package app

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"maxblog-be-demo/src/service"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	DB      *gorm.DB
	Service *service.BData
}
