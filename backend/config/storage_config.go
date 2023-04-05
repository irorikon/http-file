/*
 * @Author: iRorikon
 * @Date: 2023-04-05 02:13:51
 * @FilePath: \http-file\backend\config\storage_config.go
 */
package config

type StorageConfig struct {
	Type string `mapstructure:"type" json:"type" yaml:"type"`
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}
