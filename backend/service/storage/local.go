/*
 * @Author: iRorikon
 * @Date: 2023-04-05 21:28:52
 * @FilePath: \http-file\backend\service\storage\local.go
 */
package storage

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/model"
	"github.com/irorikon/http-file/util"
)

type LocalStorageService struct{}

func (l *LocalStorageService) List(path string) ([]*model.File, error) {
	fullPath := filepath.Join(config.CFG.Storage.Path, path)
	rawFiles, err := readDir(fullPath)
	if err != nil {
		return nil, err
	}
	var Files []*model.File
	for _, f := range rawFiles {
		if !config.CFG.Storage.ShowHidden && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		Files = append(Files, &model.File{
			Name:     f.Name(),
			Path:     filepath.Join(path, f.Name()),
			Size:     f.Size(),
			Type:     f.Mode().String(),
			IsDir:    f.IsDir(),
			Modified: f.ModTime(),
		})
	}
	return Files, nil
}

func (l *LocalStorageService) Get(path string) (*model.File, error) {
	fullPath := filepath.Join(config.CFG.Storage.Path, path)
	f, err := os.Stat(fullPath)
	if err != nil {
		return nil, err
	}
	var Fileinfo *model.File
	if f.IsDir() {
		Fileinfo.Size = 0
	} else {
		Fileinfo.Size = f.Size()
	}
	Fileinfo.Name = f.Name()
	Fileinfo.Path = path
	Fileinfo.IsDir = f.IsDir()
	Fileinfo.Modified = f.ModTime()
	return Fileinfo, nil
}

func (l *LocalStorageService) MakeDir(path string) error {
	fullPath := filepath.Join(config.CFG.Storage.Path, path)
	return os.MkdirAll(fullPath, 0755)
}

func (l *LocalStorageService) Move(srcPath, oldPath string) error {
	newFullPath := filepath.Join(config.CFG.Storage.Path, srcPath)
	oldFullPath := filepath.Join(config.CFG.Storage.Path, oldPath)
	return os.Rename(oldFullPath, newFullPath)
}

func (l *LocalStorageService) Cpoy(srcPath, dstPath string) error {
	srcFullPath := filepath.Join(config.CFG.Storage.Path, srcPath)
	dstFullPath := filepath.Join(config.CFG.Storage.Path, dstPath)
	if util.IsSubPath(srcFullPath, dstFullPath) {
		return fmt.Errorf("the destination folder is a subfolder of the source folder")
	}
	info, err := os.Stat(srcFullPath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return util.CopyDir(srcFullPath, dstFullPath)
	} else {
		return util.CopyFile(srcFullPath, dstFullPath)
	}
}

func (l *LocalStorageService) Remove(path string) error {
	fullPath := filepath.Join(config.CFG.Storage.Path, path)

	info, err := os.Stat(fullPath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return os.RemoveAll(fullPath)
	} else {
		return os.Remove(fullPath)
	}
}

// func (l *LocalStorage) Upload(path string) error {
// 	fullPath := filepath.Join(config.CFG.Storage.Path, path)
// 	out, err := os.Create(fullPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()
// }

func readDir(dirname string) ([]fs.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}
