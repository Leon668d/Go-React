package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/models"
)

var jwtKey = []byte("your_secret_key")

type UserInfo struct {
	Username string `json:"usernameVal" form:"usernameVal"`
	Password string `json:"passwordVal" form:"passwordVal"`
}

type Login struct {
	Username string `json:"usernameVal" form:"usernameVal"`
	Password string `json:"passwordVal" form:"passwordVal"`
	jwt.StandardClaims
}

// func (con UserController) Index(ctx *gin.Context) {
// 	ctx.String(200, "用户列表")
// 	//	con.success(ctx)
// }

// func (con UserController) Add(ctx *gin.Context) {
// 	ctx.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
// }

// func (con UserController) Edit(ctx *gin.Context) {
// 	username := ctx.PostForm("username")
// 	file, err := ctx.FormFile("face")
// 	// file.Filename 获取文件名称
// 	dst := path.Join("./static/upload", file.Filename)
// 	if err == nil {
// 		// 获取文件后缀名
// 		extName := path.Ext(file.Filename)
// 		allowExtMap := map[string]bool{
// 			".jpg":  true,
// 			".png":  true,
// 			".gif":  true,
// 			".jpeg": true,
// 		}
// 		_, ok := allowExtMap[extName]
// 		if ok == true {
// 			ctx.String(200, "上传的文件不合法")
// 			return
// 		}
// 		// 创建图片保存目录
// 		ctx.SaveUploadedFile(file, dst)
// 	}
// 	ctx.JSON(200, gin.H{
// 		"success":  true,
// 		"username": username,
// 		"dst":      dst,
// 	})
// 	ctx.String(200, "执行上传")
// }

func (Login) LoginCheck(c *gin.Context) {
	//	grpc服务
	// conn, err2 := grpc.Dial(":8080", grpc.WithInsecure())
	// if err2 != nil {
	// 	log.Fatalf("did not connect: %v", err2.Error())
	// }
	// defer conn.Close()

	var msg = &UserInfo{}
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "无效的 JSON 格式")
		return
	} // 将password 和 username 传给msg
	// 查看数据库中是否已经有了这个用户名，若有则进入注册页面
	var tem = &models.User{}
	err := global.DB.Where("username = ?", msg.Username).First(&tem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(200, gin.H{
			"msg": "用户名不存在",
		})
	} else if err != nil {
		c.String(200, "有其他错误")
	} else {
		if tem.Password == msg.Password {
			expirationTime := time.Now().Add(5 * time.Minute)
			claims := &Login{
				Username: msg.Username,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "生成令牌失败",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"token": tokenString,
				"msg":   "success",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "密码错误"})
		}

	}

}
