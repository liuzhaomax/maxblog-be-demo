package model

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"maxblog-be-demo/src/pb"
)

var ModelSet = wire.NewSet(
	DemoSet,
)

type Demo struct {
	gorm.Model
	Title   string `gorm:"varchar(100);not null"`
	Desc    string `gorm:"varchar(300)"`
	Content string `gorm:"varchar(30000);not null"`
	View    uint32 `gorm:"number;not null;default:0"`
	Preview string `gorm:"varchar(300)"`
	Ref     string `gorm:"varchar(300)"`
}

type Demos []Demo

func Model2PB(demo *Demo) *pb.DemoRes {
	demoRes := &pb.DemoRes{
		Id:        uint32(demo.ID),
		Title:     demo.Title,
		Desc:      demo.Desc,
		CreatedAt: demo.CreatedAt.String(),
		Content:   demo.Content,
		View:      demo.View,
		Preview:   demo.Preview,
		Ref:       demo.Ref,
	}
	return demoRes
}
