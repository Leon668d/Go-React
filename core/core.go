package core

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"main.go/global"
	"main.go/models"
)

func CreateUserTable() {
	if !global.DB.Migrator().HasTable(&models.User{}) {
		global.DB.AutoMigrate(&models.User{})
	}
}

func InitGorm() *gorm.DB {

	global.DB = MysqlConnect()
	CreateUserTable()
	return global.DB
}

func MysqlConnect() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Println("未配置mysql,取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "release" {
		// 开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s]mysql连接失败", dsn))
	}

	//	defer db.Commit().Statement.ReflectValue.Close()
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间，不能超过mysql的wait_timeout
	return db

}
