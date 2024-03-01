package global

import (
	"gorm.io/gorm"
	"main.go/config"
)

// 需要一个全局变量，用于保存配置文件，存放在 global 目录下

var (
	DB     *gorm.DB
	Config = &config.Config{}
)
