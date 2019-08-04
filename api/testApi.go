package api

import (
	"github.com/gin-gonic/gin"
)

func Ping (c *gin.Context) {
		data := map[string]interface{}{
			"lang": "理解错",
			"tag": "15217705222",
		}
		c.JSON(200, data)
}
