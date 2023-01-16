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
		u.GET("tiadd/:username", ctrl.TicketAdd)
		u.GET("coadd/:username", ctrl.CoinAdd)
		u.GET("tisub/:username", ctrl.TicketSub)
		u.GET("cosub/:username", ctrl.CoinSub)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	m := r.Group("/memos")
	{
		ctrl := controller.Controller{}
		m.GET("", ctrl.MemoIndex)
		m.GET("/showname/:username", ctrl.MemoShowname)
		m.POST("", ctrl.MemoCreate)
		m.DELETE("/:id", ctrl.MemoDelete)
	}

	return r
}
