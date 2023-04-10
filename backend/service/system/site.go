package system

import "github.com/irorikon/http-file/config"

type SiteService struct{}

// @function: GetSiteConfig
// @description: 读取配置文件
func (s *SiteService) GetSiteInfo() (config.SiteConfig, error) {
	return config.CFG.Site, nil
}

// @function: UpdateSiteConfig
// @description: 设置站点配置
func (s *SiteService) UpdateSiteConfig(site config.SiteConfig) error {
	config.CFG.Site = site
	return config.Viper.WriteConfig()
}
