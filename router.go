package main

import (
	"github.com/gin-gonic/gin"
	"go.pkg.okcoin.com/devops/httpmock/handlers"
)

func router(router *gin.Engine) {
	mockGroup := router.Group("/v1/mock")
	mockGroup.POST("/initAgent", handlers.InitAgent)
	mockGroup.PUT("/reportMachineInfo", handlers.ReportMachineInfo)
	mockGroup.PUT("/callExecCmd", handlers.ReportMachineInfo)
	mockGroup.PUT("/callDeploy", handlers.ReportMachineInfo)
	mockGroup.PUT("/monitor-files", handlers.ReportMachineInfo)
	mockGroup.POST("/monitor-process", handlers.ReportMachineInfo)
}
