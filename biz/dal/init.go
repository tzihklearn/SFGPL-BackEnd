package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysqlDsn = dsn

// DB 通过DB操控数据库
var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{
		SkipDefaultTransaction: true,                                //禁用全局事务，便于演示事务
		PrepareStmt:            true,                                //PreparedStmt 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
		Logger:                 logger.Default.LogMode(logger.Info), //日志打印配置
	})
	if err != nil {
		panic(err)
	}
}
