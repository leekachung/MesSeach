package server

import (
	"lcb-go/api"
	"lcb-go/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// cors
	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1/") 
	{
		v1.GET("ping", api.Ping)
		v1.GET("info", api.GetInfo)
	}
	return r
}
