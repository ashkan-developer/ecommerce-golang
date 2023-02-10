package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user/signup")
	incomingRoutes.POST("/user/login")
	incomingRoutes.POST("/admin/addproduct")
	incomingRoutes.GET("/users/productview")
	incomingRoutes.GET("/users/search")
}
