package router

import (
	"learn/httpserver/controller"
	"learn/httpserver/middleware"

	"github.com/gin-gonic/gin"
	// "learn/httpserver/controller"
)

func IndexRoute(route *gin.Engine) {

	// route.GET("/test", controller.Test)

	route.GET("/get", controller.Get)

	route.POST("/create", controller.Create)

	route.DELETE("/delete/:id", controller.Delete)

	route.PUT("/update/:id", controller.Update)

	route.POST("/login", controller.Login)

	route.GET("/logout", controller.Logout)

	//to authenticate with jwt
	route.GET("/auth", middleware.AuthenticateUser, middleware.ValidatePermission, controller.AuthData)

	route.GET("/session-test", middleware.AuthenticateUser, middleware.ValidatePermission, controller.SessionTest)

}
