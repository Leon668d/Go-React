package routers

import (
	"github.com/gin-gonic/gin"
	user "main.go/api/controllers"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.POST("/login", user.Login{}.LoginCheck)
		defaultRouters.GET("/register", user.Login{}.Register)
	}
}
