package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kshzz24/go-crud/controllers"
	"github.com/kshzz24/go-crud/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}
