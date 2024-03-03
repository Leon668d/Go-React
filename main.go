package main

import (
	"github.com/gin-gonic/gin"
	"main.go/api/routers"
	"main.go/core"
	"main.go/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 连接数据库
	global.DB = core.InitGorm()
	r := gin.Default()
	// 加载模板 放在配置路由上面
	//r.LoadHTMLGlob("src/*")

	// 配置session中间件

	// // 创建基于 cookie 的存储引擎， secret11111 参数是用于加密的密钥
	// store := cookie.NewStore([]byte("secret11111"))
	// // store是前面创建的存储引擎，可以替换为其他存储引擎
	// r.Use(sessions.Sessions("mysession", store))

	// 路由分组

	// CORS 中间件,前后端跨域
	cors := func(c *gin.Context) {
		// 允许特定的域进行跨域请求
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3002")
		// 允许特定的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		// 允许特定的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 允许携带身份凭证（如 Cookie）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 继续处理请求
		c.Next()
	}

	// 应用 CORS 中间件到所有路由
	r.Use(cors)

	routers.DefaultRoutersInit(r)
	r.Run(":8080")

}
