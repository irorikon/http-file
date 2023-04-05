/*
 * @Author: iRorikon
 * @Date: 2023-04-04 16:32:08
 * @FilePath: \http-file\backend\config\config.go
 */
package config

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	ConfigFile string
	Version    bool
	CFG        *Config
	Log        *zap.Logger
	BlackCache local_cache.Cache
	Viper      *viper.Viper
)

const (
	ConfigEnv         = "HTTP_FILE_CONFIG"
	DefaultConfigFile = "./config.yaml"
)

var SlicesMap = make(map[string][]string)

const (
	TextTypes  = "text_types"
	AudioTypes = "audio_types"
	VideoTypes = "video_types"
	ImageTypes = "image_types"
)

const (
	UNKNOWN = iota
	FOLDER
	//OFFICE
	VIDEO
	AUDIO
	TEXT
	IMAGE
)

type Config struct {
	System  SystemConfig  `mapstructure:"system" json:"system" yaml:"system"`
	Server  ServerConfig  `mapstructure:"server" json:"server" yaml:"server"`
	Storage StorageConfig `mapstructure:"storage" json:"storage" yaml:"storage"`
	Zap     ZapConfig     `mapstructure:"zap" json:"zap" yaml:"zap"`
	CORS    CORSConfig    `mapstructure:"cors" json:"cors" yaml:"cors"`
}
