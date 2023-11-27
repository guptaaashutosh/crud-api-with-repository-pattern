package router

import (
	"learn/httpserver/controller"

	"github.com/gin-gonic/gin"
	// "learn/httpserver/controller"
)

func IndexRoute(route *gin.Engine) {

	// route.GET("/test", controller.Test)

	route.GET("/get", controller.Get)

	route.POST("/create", controller.Create)

	route.DELETE("/delete/:id", controller.Delete)

	route.PUT("/update/:id", controller.Update)



}
