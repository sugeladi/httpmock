package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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

func ReportMachineInfo(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("readAll err:%v\n", err.Error())
	}
	fmt.Printf("body:%v \n", string(body))
	c.JSON(http.StatusOK, gin.H{
		"requestId": "b92572b5-fc6c-4a7d-baf9-5c14ecfde781",
		"code":      0,
		"msg":       "",
		"detailMsg": "",
		"success":   true,
		"data":      "",
	})
	return
}
