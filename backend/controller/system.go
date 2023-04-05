/*
 * @Author: iRorikon
 * @Date: 2023-04-05 02:04:58
 * @FilePath: \http-file\backend\controller\system.go
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/model/response"
	"go.uber.org/zap"
)

type SystemController struct{}

func (sc *SystemController) GetSystemConfig(c *gin.Context) {
	cfg, err := systemService.GetSystemInfo()
	if err != nil {
		config.Log.Error("获取配置文件失败", zap.String("err", err.Error()))
		response.FailWithMessage("获取配置文件失败", c)
		return
	}
	response.OKWithDetailed(cfg, "获取配置文件成功", c)

}
