package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/api/controllers/login"
	createuser "main.go/api/controllers/register"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.POST("/login", login.UserController{}.LoginCheck)
		defaultRouters.GET("/register", createuser.UserMsg{}.CreateUser)
	}
}
