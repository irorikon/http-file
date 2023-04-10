/*
 * @Author: iRorikon
 * @Date: 2023-04-07 21:45:49
 * @FilePath: \http-file\backend\controller\storage\storage.go
 */
package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/model/response"
	"go.uber.org/zap"
)

type StorageController struct{}

// @function: GetStorageConfig
// @description: 读取存储配置文件
func (sc *StorageController) GetStorageConfig(c *gin.Context) {
	cfg, err := storageService.GetStorageConfig()
	if err != nil {
		config.Log.Error("Failed to get storage configuration", zap.String("err", err.Error()))
		response.FailWithMessage("Failed to get storage configuration", c)
		return
	}
	response.OKWithDetailed(cfg, "Get storage configuration successfully", c)
}

// @function: UpdateStorageConfig
// @description: 更新存储配置文件
func (sc *StorageController) UpdateStorageConfig(c *gin.Context) {
	var cfg config.StorageConfig
	if err := c.ShouldBindJSON(&cfg); err != nil {
		response.FailWithMessage("Parameter error", c)
		return
	}
	if err := storageService.UpdateStorageConfig(cfg); err != nil {
		response.FailWithMessage("Configuration settings failed", c)
		return
	}
	response.OKWithMessage("Configuration setup successful", c)
}
