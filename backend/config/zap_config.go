/*
 * @Author: iRorikon
 * @Date: 2023-04-04 17:56:02
 * @FilePath: \http-file\backend\config\zap_config.go
 */
package config

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type ZapConfig struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Director      string `mapstructure:"director" json:"director" yaml:"director"`
	EncodeLevel   string `mapstructure:"encode_level" json:"encode_level" yaml:"encode_level"`
	StacktraceKey string `mapstructure:"stacktrace_key" json:"stacktrace_key" yaml:"stacktrace_key"`
	MaxAge        int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	ShowLine      bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	LogToConsole  bool   `mapstructure:"log_to_console" json:"log_to_console" yaml:"log_to_console"`
}

func (z *ZapConfig) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 根据字符串转化为 zapcore.Level
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *ZapConfig) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
