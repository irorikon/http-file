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

func DefaultConfig() *Config {
	return &Config{
		System: SystemConfig{
			Title:    "HTTP-FILE",
			MusicIMG: "https://img.oez.cc/2020/12/19/0f8b57866bdb5.gif",
			AutoPlay: true,
			Sort:     "name-descend",
			Preview: PreviewConfig{
				Text: []string{
					"txt",
					"htm",
					"html",
					"xml",
					"java",
					"properties",
					"sql",
					"js",
					"md",
					"json",
					"conf",
					"ini",
					"vue",
					"php",
					"py",
					"bat",
					"gitignore",
					"yml",
					"go",
					"sh",
					"c",
					"cpp",
					"h",
					"hpp",
				},
			},
		},
		Server: ServerConfig{
			Address:  "0.0.0.0",
			Port:     "5244",
			Password: "",
		},
		Storage: StorageConfig{
			Type:       "local",
			ShowHidden: false,
			Path:       "./",
		},
		Zap: ZapConfig{
			Level:         "debug",
			Prefix:        "[HTTP-FILE]",
			FilePrefix:    "http-file",
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
