package server

import (
	"github.com/gin-gonic/gin"

	"login-api/controller"
)

// サーバー起動
func Init() {
	r := router()
	r.Run(":8081")
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := controller.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/showid/:id", ctrl.Showid)
		u.GET("/showname/:username", ctrl.Showname)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}
