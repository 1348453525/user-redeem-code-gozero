package svc

import (
	"log"
	"os"
	"time"

	"github.com/1348453525/user-redeem-code-gozero/user-rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB // 注入GORM实例
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 创建GORM实例
	db, err := NewGormDB(c.Mysql.DataSource)
	if err != nil {
		logx.Errorw("数据库连接失败", logx.Field("err", err))
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}

// NewGormDB 初始化GORM连接
func NewGormDB(dsn string) (*gorm.DB, error) {
	// 配置GORM日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色打印
			// IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound
		},
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	// 连接池配置
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)           // 空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最长存活时间

	return db, nil
}
