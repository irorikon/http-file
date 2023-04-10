/*
 * @Author: iRorikon
 * @Date: 2023-04-05 21:28:52
 * @FilePath: \http-file\backend\service\storage\storage.go
 */
package storage

import (
	"github.com/irorikon/http-file/config"
)

type StorageService struct{}

// @function: GetStorageConfig
// @description: 读取配置文件
func (s *StorageService) GetStorageConfig() (*config.StorageConfig, error) {
	return &config.CFG.Storage, nil
}

// @function: UpdateStorageConfig
// @description: 设置配置文件
func (s *StorageService) UpdateStorageConfig(cfg config.StorageConfig) error {
	config.CFG.Storage = cfg
	return config.Viper.WriteConfig()
}
