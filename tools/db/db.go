package db

import (
	"log"
	"os"
	"time"

	"github.com/5-say/go-tools/tools/change"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config ..
type Config struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	LogLevel      int // 日志级别 1.Silent 2.Error 3.Warn 4.Info
	SlowThreshold int // 慢 SQL 阈值（秒）
}

// LogWriter ..
func LogWriter() logger.Writer {
	return log.New(os.Stdout, "\r\n", log.LstdFlags) // io writer
}

// Open ..
func Open(config Config, logWriter logger.Writer) *gorm.DB {
	dsn := config.User + ":" + config.Password + "@(" + config.Host + ":" + change.ToString(config.Port) + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	logLevel := map[int]logger.LogLevel{1: logger.Silent, 2: logger.Error, 3: logger.Warn, 4: logger.Info}[config.LogLevel]
	newLogger := logger.New(
		logWriter, // io writer
		logger.Config{
			Colorful:      false,                                             // 彩色打印
			LogLevel:      logLevel,                                          // 日志级别
			SlowThreshold: time.Second * time.Duration(config.SlowThreshold), // 慢 SQL 阈值
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 是否跳过 “根据当前 MySQL 版本自动配置”
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭迁移时的自动外键关联
		SkipDefaultTransaction:                   true, // 关闭自动全局事务
		Logger:                                   newLogger,
	})

	if err != nil {
		log.Fatalf("Fail to db logger: %v\n", err)
	}

	return db
}
