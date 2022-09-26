package service

import (
	"context"
	"github.com/google/wire"
	"maxblog-be-demo/internal/core"
	"maxblog-be-demo/src/model"
	"maxblog-be-demo/src/pb"
)

var DemoSet = wire.NewSet(wire.Struct(new(BDemo), "*"))

type BDemo struct {
	MDemo   *model.MDemo
	Tx      *core.Trans
	ILogger core.ILogger
}

func (bDemo *BDemo) GetDemos(ctx context.Context, req *pb.CountRequest) (*pb.DemosRes, error) {
	var demos model.Demos
	var pageSize uint32 = 6
	var offset = pageSize * req.Count
	err := bDemo.MDemo.QueryDemos(pageSize, offset, &demos)
	if err != nil {
		bDemo.ILogger.LogFailure(core.GetFuncName(), err)
		return nil, err
	}
	var res pb.DemosRes
	for _, item := range demos {
		res.DataArr = append(res.DataArr, model.Model2PB(&item))
	}
	bDemo.ILogger.LogSuccess(core.GetFuncName())
	return &res, nil
}

func (bDemo *BDemo) GetDemoById(ctx context.Context, req *pb.IdRequest) (*pb.DemoRes, error) {
	var data model.Demo
	err := bDemo.Tx.ExecTrans(ctx, func(ctx context.Context) error {
		err := bDemo.MDemo.QueryDemoById(req, &data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		bDemo.ILogger.LogFailure(core.GetFuncName(), err)
		return nil, err
	}
	res := model.Model2PB(&data)
	bDemo.ILogger.LogSuccess(core.GetFuncName())
	return res, nil
}
