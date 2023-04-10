/*
 * @Author: iRorikon
 * @Date: 2023-04-05 14:34:06
 * @FilePath: \http-file\backend\service\system\system.go
 */
package system

import (
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/util"
	"go.uber.org/zap"
)

type SystemService struct{}

// @function: GetSystemConfig
// @description: 读取配置文件
func (s *SystemService) GetSystemConfig() (config.SystemConfig, error) {
	return config.CFG.System, nil
}

// @function: SetSystemConfig
// @description: 设置配置文件
func (s *SystemService) UpdateSystemConfig(system config.SystemConfig) error {
	config.CFG.System = system
	return config.Viper.WriteConfig()
}

// @function: GetSystemState
// @description: 获取服务器信息
func (s *SystemService) GetSystemState() (server *util.Server, err error) {
	var svr util.Server
	svr.OS = util.InitOS()
	if svr.CPU, err = util.InitCPU(); err != nil {
		config.Log.Error("获取CPU信息失败", zap.String("err", err.Error()))
		return &svr, err
	}
	if svr.RAM, err = util.InitRAM(); err != nil {
		config.Log.Error("获取RAM信息失败", zap.String("err", err.Error()))
		return &svr, err
	}
	if svr.Disk, err = util.InitDisk(); err != nil {
		config.Log.Error("获取Disk信息失败", zap.String("err", err.Error()))
		return &svr, err
	}
	return &svr, nil
}
