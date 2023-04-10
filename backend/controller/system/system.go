/*
 * @Author: iRorikon
 * @Date: 2023-04-05 02:04:58
 * @FilePath: \http-file\backend\controller\system\system.go
 */
package system

import (
	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/model/response"
	"go.uber.org/zap"
)

type SystemController struct{}

// @function: GetSystemConfig
// @description: 读取配置文件
func (sc *SystemController) GetSystemConfig(c *gin.Context) {
	cfg, err := systemService.GetSystemConfig()
	if err != nil {
		config.Log.Error("获取配置文件失败", zap.String("err", err.Error()))
		response.FailWithMessage("获取配置文件失败", c)
		return
	}
	response.OKWithDetailed(cfg, "获取配置文件成功", c)

}

// @function: UpdateSystemConfig
// @description: 设置配置文件
func (sc *SystemController) UpdateSystemConfig(c *gin.Context) {
	var cfg config.SystemConfig
	if err := c.ShouldBindJSON(&cfg); err != nil {
		response.FailWithMessage("Parameter error", c)
		return
	}
	if err := systemService.UpdateSystemConfig(cfg); err != nil {
		response.FailWithMessage("Configuration settings failed", c)
		return
	}
	response.OKWithMessage("Configuration setup successful", c)
}

// @function: GetSystemState
// @description: 获取系统运行状态
func (sc *SystemController) GetSystemState(c *gin.Context) {
	info, err := systemService.GetSystemState()
	if err != nil {
		config.Log.Error("获取系统运行状态失败", zap.String("err", err.Error()))
		response.FailWithMessage("获取系统运行状态失败", c)
		return
	}
	response.OKWithDetailed(info, "获取系统运行状态成功", c)
}
