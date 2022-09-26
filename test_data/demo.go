package test_data

import (
	"fmt"
	"gorm.io/gorm"
	"maxblog-be-demo/src/model"
)

func CreateDemoData(db *gorm.DB) {
	var data model.Demo
	result := db.First(&data)
	if result.RowsAffected == 0 {
		for i := 0; i < 18; i++ {
			var demo model.Demo
			demo.Title = fmt.Sprintf("小刘鸭%d", i)
			demo.Desc = "我是一只丑小鸭"
			demo.Content = "for i, abc = range abcs{}"
			demo.View = uint32(i)
			demo.Preview = "http://10.192.0.5:9202/static/demo/duck_liu/02.jpg"
			demo.Ref = "https://www.baidu.com"
			db.Create(&demo)
		}
	}
}
