package main

import (
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	myflag "go.pkg.okcoin.com/devops/httpmock/internal/flag"
	"os"
)

//程序信息
const (
	SERVICE_NAME = "httpmock"
	VERSION      = "1.0.0"
)

func main() {

	//读取程序运行参数并验证
	isVersion, flagInit, err := myflag.InitFlags()
	if isVersion {
		fmt.Fprintf(os.Stdout, "%s\t%s\n", SERVICE_NAME, VERSION)
		return
	} else if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}

	//gin设置
	gin.SetMode(gin.ReleaseMode)

	//路由设置
	engine := gin.Default()
	engine.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	})
	router(engine)

	//服务启动
	fmt.Printf("%s Service start ok,address=%s:%s \n", SERVICE_NAME, flagInit.ServiceAddress, flagInit.ServicePort)
	ginpprof.Wrapper(engine)
	if err := engine.Run(fmt.Sprintf("%s:%s", flagInit.ServiceAddress, flagInit.ServicePort)); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}
