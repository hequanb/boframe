package mysqlI

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	
	"boframe/settings"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init(config *settings.MYSQLConfig) (err error) {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
	
	// TODO: gorm整合zap
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 10 * time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info,      // Log level
		},
	)
	
	db, err = gorm.Open(mysql.Open(datasourceName), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		FullSaveAssociations:                     false,
		Logger:                                   newLogger,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})
	if err != nil {
		zap.L().Error("connect DB failed \n", zap.Error(err))
		return err
	}
	
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("connect DB failed \n", zap.Error(err))
		return err
	}
	
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	
	// ping你一下
	return sqlDB.Ping()
}

// TODO: concurrency safe
func DB() *gorm.DB {
	return db
}

func Close() {
}

func IsErrNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
