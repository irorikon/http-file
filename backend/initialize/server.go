/*
 * @Author: iRorikon
 * @Date: 2023-04-05 03:00:39
 * @FilePath: \http-file\backend\initialize\server.go
 */
package initialize

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/router"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
	ListenAndServeTLS(string, string) error
}

func RunServer() {
	if config.CFG.Zap.Level == "info" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	Router := router.Routers()
	// Router.Static("/", "./templates")
	address := fmt.Sprintf("%v:%v", config.CFG.System.Address, config.CFG.System.Port)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted

	time.Sleep(10 * time.Microsecond)
	config.Log.Info("server run success on ", zap.String("address", address))
	if config.CFG.System.TLS {
		config.Log.Error(s.ListenAndServeTLS(config.CFG.System.CertFile, config.CFG.System.KeyFile).Error())
	}
	config.Log.Error(s.ListenAndServe().Error())
}
