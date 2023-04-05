package service

import (
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/model/request"
	"github.com/irorikon/http-file/util"
	"go.uber.org/zap"
)

type SystemService struct{}

// @function: GetSystemConfig
// @description: 读取配置文件
func (s *SystemService) GetSystemInfo() (config.SystemConfig, error) {
	return config.CFG.System, nil
}

// @function: SetSystemConfig
// @description: 设置配置文件
func (s *SystemService) SetSystemConfig(system request.SystemReq) error {
	cs := util.StructToMap(system.System)
	for k, v := range cs {
		config.Viper.Set(k, v)
	}
	err := config.Viper.WriteConfig()
	return err
}

// @function: GetServerInfo
// @description: 获取服务器信息
func (s *SystemService) GetServerInfo() (server *util.Server, err error) {
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
