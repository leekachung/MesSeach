package server

import (
	"lcb-go/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1/") 
	{
		v1.GET("ping", api.Ping)
	}
	return r
}
