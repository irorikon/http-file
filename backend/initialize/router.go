/*
 * @Author: iRorikon
 * @Date: 2023-04-05 03:10:34
 * @FilePath: \http-file\backend\initialize\router.go
 */
package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/controller"
	"github.com/irorikon/http-file/middleware"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.Default()

	Router.LoadHTMLGlob("./templates/*.html") // npm打包成dist的路径
	Router.Static("/favicon.ico", "./templates/favicon.ico")
	Router.Static("/assets", "./templates/assets")   // dist里面的静态资源
	Router.StaticFile("/", "./templates/index.html") // 前端网页入口页面

	// TLS
	// 如果需要使用 https 请打开此中间件
	// 然后前往 core/server.go 将启动模式更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	if config.CFG.Server.TLS {
		if config.CFG.Server.KeyFile != "" && config.CFG.Server.CertFile != "" {
			Router.Use(middleware.LoadTLS())
		} else {
			config.CFG.Server.TLS = false
		}
	}
	// 跨域，如需跨域可以打开下面的注释
	// 直接放行全部跨域请求
	// Router.Use(middleware.CORS())
	// 按照配置的规则放行跨域请求
	Router.Use(middleware.CORSByRules())
	config.Log.Info("use middleware cors")

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		PublicGroup.GET("/info", controller.ControllerAPP.SystemController.GetSystemConfig)
		// PublicGroup.POST("/path", controller.ControllerAPP.GetFileList)
	}

	config.Log.Info("router init success")
	return Router
}
