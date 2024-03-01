package createuser

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/models"
)

type UserMsg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (UserMsg) CreateUser(c *gin.Context) {
	var msg = &UserMsg{}
	c.ShouldBindJSON(&msg) // 将password 和 username 传给msg
	// 查看数据库中是否已经有了这个用户名，若有则进入注册页面
	var tem = &models.User{}
	err := global.DB.Where("username = ?", msg.Username).First(&tem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // 该用户名未被使用
		var u = &models.User{
			Username: msg.Username,
			Password: msg.Password,
		}
		global.DB.Create(&u) // 将该用户用户名以及密码写入数据库
		c.JSON(200, gin.H{
			"ID":       tem.ID,
			"username": tem.Username,
			"msg":      "success",
		})
	} else if err != nil {
		// 其他错误
	} else { // 用户名已被使用
		c.JSON(200, gin.H{
			"username": tem.Username,
			"msg":      "fail",
		})
	}
}
