package model

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"maxblog-be-demo/internal/core"
	"maxblog-be-demo/src/pb"
)

var DemoSet = wire.NewSet(wire.Struct(new(MDemo), "*"))

type MDemo struct {
	DB *gorm.DB
}

func (mDemo *MDemo) QueryDemoById(req *pb.IdRequest, data *Demo) error {
	result := mDemo.DB.First(data, req.Id)
	if result.RowsAffected == 0 {
		return core.FormatError(803, nil)
	}
	return nil
}

func (mDemo *MDemo) QueryDemos(total uint32, demos *Demos) error {
	result := mDemo.DB.Find(&demos).Limit(int(total))
	if result.RowsAffected == 0 {
		return core.FormatError(803, nil)
	}
	return nil
}
