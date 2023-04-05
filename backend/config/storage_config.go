/*
 * @Author: iRorikon
 * @Date: 2023-04-05 02:13:51
 * @FilePath: \http-file\backend\config\storage_config.go
 */
package config

type StorageConfig struct {
	Type       string `mapstructure:"type" json:"type" yaml:"type"`
	ShowHidden bool   `mapstructure:"show_hidden" json:"show_hidden" yaml:"show_hidden"`
	Path       string `mapstructure:"path" json:"path" yaml:"path"`
}
