package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAgent(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"requestId": "b92572b5-fc6c-4a7d-baf9-5c14ecfde781",
		"code":      0,
		"msg":       "",
		"detailMsg": "",
		"success":   true,
		"data": gin.H{
			"agentId": "10000",
		},
	})
	return
}
