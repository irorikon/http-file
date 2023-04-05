/*
 * @Author: iRorikon
 * @Date: 2023-04-04 17:15:36
 * @FilePath: \backend\util\file_operations.go
 */
package util

import (
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/irorikon/http-file/config"
)

// @function: FileMove
// @description: 文件移动供外部调用
// @param: src string, dst string(src: 源位置,绝对路径or相对路径, dst: 目标位置,绝对路径or相对路径,必须为文件夹)
// @return: err error
func FileMove(src string, dst string) (err error) {
	if dst == "" {
		return nil
	}
	src, err = filepath.Abs(src)
	if err != nil {
		return err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}
	revoke := false
	dir := filepath.Dir(dst)
Redirect:
	_, err = os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0o755)
		if err != nil {
			return err
		}
		if !revoke {
			revoke = true
			goto Redirect
		}
	}
	return os.Rename(src, dst)
}

// @function: DeLFile
// @description: 删除
func DeLFile(filePath string) error {
	return os.RemoveAll(filePath)
}

// @function: TrimSpace
// @description: 去除结构体空格
func TrimSpace(target interface{}) {
	t := reflect.TypeOf(target)
	if t.Kind() != reflect.Ptr {
		return
	}
	t = t.Elem()
	v := reflect.ValueOf(target).Elem()
	for i := 0; i < t.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			v.Field(i).SetString(strings.TrimSpace(v.Field(i).String()))
		}
	}
}

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	fi, err := os.Lstat(path)
	if err == nil {
		return !fi.IsDir()
	}
	return !os.IsNotExist(err)
}

// GetFileType get file type
func GetFileType(filename string) int {
	ext := strings.ToLower(Ext(filename))
	if SliceContains(config.SlicesMap[config.AudioTypes], ext) {
		return config.AUDIO
	}
	if SliceContains(config.SlicesMap[config.VideoTypes], ext) {
		return config.VIDEO
	}
	if SliceContains(config.SlicesMap[config.ImageTypes], ext) {
		return config.IMAGE
	}
	if SliceContains(config.SlicesMap[config.TextTypes], ext) {
		return config.TEXT
	}
	return config.UNKNOWN
}

func Ext(p string) string {
	ext := path.Ext(p)
	if strings.HasPrefix(ext, ".") {
		return ext[1:]
	}
	return ext
}
