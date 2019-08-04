package api

import (
	"lcb-go/service"
	"lcb-go/serializer"
	"github.com/gin-gonic/gin"
)

func GetInfo (c *gin.Context) {
	var service service.SearchService
	if c.ShouldBind(&service) == nil {
		if info := service.Search(); info == nil {
			c.JSON(200, serializer.Response{
				Status: 40032,
				Msg: "Not Data",
			})
		} else {
			c.JSON(200, serializer.Response{
				Status: 200,
				Data: info,
			})
		}
	} else {	
		c.JSON(200, serializer.Response{
			Status: 40002,
			Msg: "error",
		})
	}
}
