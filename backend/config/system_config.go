/*
 * @Author: iRorikon
 * @Date: 2023-04-05 02:11:43
 * @FilePath: \http-file\backend\config\system_config.go
 */
package config

type SystemConfig struct {
	Address             string `mapstructure:"address" json:"address" yaml:"address"`
	Port                string `mapstructure:"port" json:"port" yaml:"port"`
	TLS                 bool   `mapstructure:"tls" json:"tls" yaml:"tls"`
	CertFile            string `mapstructure:"cert_file" json:"cert_file" yaml:"cert_file"`
	KeyFile             string `mapstructure:"key_file" json:"key_file" yaml:"key_file"`
	Prefix              string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Search              bool   `mapstructure:"search" json:"search" yaml:"search"`
	Download            bool   `mapstructure:"download" json:"download" yaml:"download"`
	Static              string `mapstructure:"static" json:"static" yaml:"static"`
	SiteURL             string `mapstructure:"site_url" json:"site_url" yaml:"site_url"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	FileDownloadTimeout int64  `mapstructure:"file_download_timeout" json:"file_download_timeout" yaml:"file_download_timeout"`
}
