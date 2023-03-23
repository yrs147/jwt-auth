package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/yrs147/jwt-auth/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
	incomingRoutes.POST("users/logout", controller.Logout())

}
