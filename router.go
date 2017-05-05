package main

import (
	"github.com/gin-gonic/gin"
	"go.pkg.okcoin.com/devops/httpmock/handlers"
)

func router(router *gin.Engine) {
	mockGroup := router.Group("/v1/mock")
	mockGroup.POST("/initAgent", handlers.InitAgent)
	mockGroup.POST("/reportMachineInfo", handlers.ReportMachineInfo)
}
