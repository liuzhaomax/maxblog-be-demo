package conf

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"maxblog-be-demo/internal/core"
	"maxblog-be-demo/src/model"
	"maxblog-be-demo/test_data"
	"os"
	"strings"
	"time"
)

func (cfg *Config) NewDB() (*gorm.DB, func(), error) {
	fileName := "golog.txt"
	logFilePath, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	logg := logger.New()
	logg.Out = logFilePath
	logg.SetLevel(logger.InfoLevel)
	logg.SetFormatter(&logger.TextFormatter{ForceColors: cfg.Logger.Color})
	logg.SetOutput(io.MultiWriter(logFilePath, os.Stdout))
	gLogger := gormLogger.New(
		logg,
		gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  cfg.Logger.Color,
		},
	)
	logger.Info(fmt.Sprintf("数据库种类: %s", cfg.DB.Type))
	cfg.DB.DSN = cfg.Mysql.DSN()
	db, err := gorm.Open(mysql.Open(cfg.DB.DSN), &gorm.Config{
		Logger: gLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	if cfg.DB.Debug {
		db = db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	clean := func() {
		err = sqlDB.Close()
		if err != nil {
			logger.WithFields(logger.Fields{
				"失败方法": core.GetFuncName(),
			}).Error(core.FormatError(800, err).Error())
		}
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, clean, err
	}
	sqlDB.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DB.MaxLifetime) * time.Second)
	return db, clean, err
}

func (cfg *Config) AutoMigrate(db *gorm.DB) error {
	dbType := strings.ToLower(cfg.DB.Type)
	if dbType == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	err := db.AutoMigrate(new(model.Demo))
	if err != nil {
		return err
	}
	// 测试数据
	test_data.CreateDemoData(db)
	return nil
}
