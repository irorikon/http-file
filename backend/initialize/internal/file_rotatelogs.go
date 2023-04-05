/*
 * @Author: iRorikon
 * @Date: 2023-04-04 18:02:15
 * @FilePath: \http-file\backend\initialize\internal\file_rotatelogs.go
 */
package internal

import (
	"os"
	"path"
	"time"

	"github.com/irorikon/http-file/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

type fileRotatelogs struct{}

var FileRotatelogs = new(fileRotatelogs)

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotatelogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(config.CFG.Zap.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(config.CFG.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if config.CFG.Zap.LogToConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
