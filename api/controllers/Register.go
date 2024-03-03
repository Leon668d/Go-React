package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/models"
)

func (Login) Register(c *gin.Context) {
	var msg = &Login{}
	c.ShouldBindJSON(&msg) // 将password 和 username 传给msg
	var tem = &models.User{}
	err := global.DB.Where("username = ?", msg.Username).First(&tem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // 该用户名未被使用
		var u = &models.User{
			Username: msg.Username,
			Password: msg.Password,
		}
		global.DB.Create(&u) // 将该用户用户名以及密码写入数据库
		c.JSON(200, gin.H{
			"username": tem.Username,
			"msg":      "成功注册",
		})
	} else if err != nil {
		// 其他错误
	} else { // 用户名已被使用
		c.JSON(200, gin.H{
			"username": msg.Username,
			"msg":      "该用户名已被使用",
		})
	}
}
