package initialize

import (
	"fmt"
	"os"

	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/initialize/internal"
	"github.com/irorikon/http-file/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := util.PathExists(config.CFG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.CFG.Zap.Director)
		_ = os.Mkdir(config.CFG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if config.CFG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
