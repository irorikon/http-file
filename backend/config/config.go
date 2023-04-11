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
	Site    SiteConfig    `mapstructure:"site" json:"site" yaml:"site"`
	JWT     JWTConfig     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Storage StorageConfig `mapstructure:"storage" json:"storage" yaml:"storage"`
	Zap     ZapConfig     `mapstructure:"zap" json:"zap" yaml:"zap"`
	CORS    CORSConfig    `mapstructure:"cors" json:"cors" yaml:"cors"`
}

func DefaultConfig() *Config {
	return &Config{
		Site: SiteConfig{
			Title: "HTTP-FILE",
			Sort:  "name-descend",
		},
		System: SystemConfig{
			Address:             "0.0.0.0",
			Port:                "5244",
			Password:            "",
			FileDownloadTimeout: 86400,
		},
		Storage: StorageConfig{
			Type:       "local",
			ShowHidden: false,
			Path:       "./",
		},
		Zap: ZapConfig{
			Level:         "debug",
			Prefix:        "[HTTP-FILE]",
			Format:        "json",
			Director:      "logs",
			EncodeLevel:   "LowercaseLevelEncoder",
			StacktraceKey: "stacktrace",
			MaxAge:        3600,
			ShowLine:      true,
			LogToConsole:  true,
		},
		CORS: CORSConfig{
			Mode: "allow-all",
		},
	}
}
