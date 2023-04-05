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
	Router := Routers()
	// Router.Static("/", "./templates")
	address := fmt.Sprintf("%v:%v", config.CFG.Server.Address, config.CFG.Server.Port)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted

	time.Sleep(10 * time.Microsecond)
	config.Log.Info("server run success on ", zap.String("address", address))
	if config.CFG.Server.TLS {
		config.Log.Error(s.ListenAndServeTLS(config.CFG.Server.CertFile, config.CFG.Server.KeyFile).Error())
	}
	config.Log.Error(s.ListenAndServe().Error())
}
