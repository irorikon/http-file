package util

import (
	"errors"
	"os"

	"github.com/irorikon/http-file/config"
	"go.uber.org/zap"
)

// @function: PathExists
// @description: 文件目录是否存在
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// @function: CreateDir
// @description: 批量创建文件夹
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			config.Log.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				config.Log.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
