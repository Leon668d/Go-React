package login

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/models"
)

// // 签名密钥
// const sign_key = "mysecret"

// // 随机字符串
// var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func randStr(str_len int) string {
// 	rand_bytes := make([]rune, str_len)
// 	for i := range rand_bytes {
// 		rand_bytes[i] = letters[rand.Intn(len(letters))]
// 	}
// 	return string(rand_bytes)
// }

// func generateTokenUsingHs256() (string, error) {
// 	claim := Login{
// 		ID:       000001,
// 		Username: "Tom",
// 		Password: "read_user_info",
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			Issuer:    "Auth_Server",                                   // 签发者
// 			Subject:   "Tom",                                           // 签发对象
// 			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},      //签发受众
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
// 			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
// 			ID:        randStr(10),                                     // wt ID, 类似于盐值
// 		},
// 	}
// 	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(sign_key))
// 	return token, err
// }

// func parseTokenHs256(token_string string) (*Login, error) {
// 	token, err := jwt.ParseWithClaims(token_string, &Login{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(sign_key), nil //返回签名密钥
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, errors.New("claim invalid")
// 	}

// 	claims, ok := token.Claims.(*Login)
// 	if !ok {
// 		return nil, errors.New("invalid claim type")
// 	}

// 	return claims, nil
// }

type Login struct {
	ID                   uint   `gorm:"primarykey"`
	Username             string `json:"usernameVal" form:"usernameVal"`
	Password             string `json:"passwordVal" form:"passwordVal"`
	jwt.RegisteredClaims        // 标准的 jwt payload。
}

type UserController struct {
	Login
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

func (UserController) LoginCheck(c *gin.Context) {
	var msg = &Login{}
	c.ShouldBindJSON(&msg) // 将password 和 username 传给msg
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
		// 找到记录
		var tId = tem.ID
		if msg.Password == tem.Password {
			c.JSON(200, gin.H{
				"ID":       tId,
				"username": tem.Username,
				"msg":      "success",
			})
		} else {
			c.JSON(200, gin.H{
				"msg": "密码错误",
			})
		}
	}

}
