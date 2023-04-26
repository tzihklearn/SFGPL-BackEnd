package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 配置数据库连接
const dsn = "service:service@tcp(43.142.146.75:3306)/sfgpl?charset=utf8mb4&parseTime=True&loc=Local"

// 通过DB操控数据库
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,                                //禁用全局事务，提升性能
		PrepareStmt:            true,                                //PreparedStmt 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
		Logger:                 logger.Default.LogMode(logger.Info), //日志打印配置
	})
	if err != nil {
		panic(err)
	}
}
