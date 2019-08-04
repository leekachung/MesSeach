package api

import (
	"lcb-go/serializer"
	"github.com/gin-gonic/gin"
)

func Ping (c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "Pong",
	})
}
