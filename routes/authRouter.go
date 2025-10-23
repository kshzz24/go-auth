package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kshzz24/go-crud/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("user/signup", controllers.Signup())
	incomingRoutes.POST("user/login", controllers.Login())
}
