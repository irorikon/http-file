package system

import (
	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/model/response"
	"go.uber.org/zap"
)

type SiteController struct{}

// @function: GetSiteConfig
// @description: 读取配置文件
func (sc *SiteController) GetSiteConfig(c *gin.Context) {
	cfg, err := siteService.GetSiteInfo()
	if err != nil {
		config.Log.Error("Failed to get configuration", zap.String("err", err.Error()))
		response.FailWithMessage("Failed to get configuration", c)
		return
	}
	response.OKWithDetailed(cfg, "Get the configuration successfully", c)
}

// @function: UpdateSiteConfig
// @description: 更新配置文件
func (sc *SiteController) UpdateSiteConfig(c *gin.Context) {
	var siteConfig config.SiteConfig
	if err := c.ShouldBindJSON(&siteConfig); err != nil {
		config.Log.Error("Failed to update configuration", zap.String("err", err.Error()))
		response.FailWithMessage("Failed to update configuration", c)
		return
	}
	if err := siteService.UpdateSiteConfig(siteConfig); err != nil {
		config.Log.Error("Failed to update configuration", zap.String("err", err.Error()))
		response.FailWithMessage("Failed to update configuration", c)
		return
	}
	response.OK(c)
}
