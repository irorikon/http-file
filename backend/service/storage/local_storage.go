/*
 * @Author: iRorikon
 * @Date: 2023-04-07 10:08:03
 * @FilePath: \http-file\backend\service\storage\local_storage.go
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

func (l *LocalStorageService) GetFileList(path string) ([]*model.File, error) {
	fullPath := filepath.Join(config.CFG.Storage.Path, path)
	rawFiles, err := readDir(fullPath)
	if err != nil {
		return nil, err
	}
	var Files []*model.File
	for _, f := range rawFiles {
		var fileType string
		var fileMIME string
		var ext string
		if !config.CFG.Storage.ShowHidden && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		if !f.IsDir() {
			fullFilePath := filepath.Join(fullPath, f.Name())
			fileType, fileMIME, err = util.FileType(fullFilePath)
			if err != nil {
				return nil, err
			}
			ext = util.FileExtension(f.Name())
		}
		Files = append(Files, &model.File{
			Name:     f.Name(),
			Path:     filepath.Join(path, f.Name()),
			Size:     f.Size(),
			Type:     fileType,
			MIME:     fileMIME,
			Ext:      ext,
			IsDir:    f.IsDir(),
			Modified: f.ModTime(),
		})
	}
	return Files, nil
}

func (l *LocalStorageService) GetFileInfo(path string) (*model.File, error) {
	fullPath := filepath.Join(config.CFG.Storage.Path, path)
	f, err := os.Stat(fullPath)
	if err != nil {
		return nil, err
	}
	var ext string
	var fileType string
	var fileMIME string
	var size int64
	if f.IsDir() {
		size = 0
	} else {
		size = f.Size()
		ext = util.FileExtension(fullPath)
		fileType, fileMIME, err = util.FileType(fullPath)
		if err != nil {
			return nil, err
		}
	}
	return &model.File{
		Name:     f.Name(),
		Path:     path,
		Size:     size,
		Type:     fileType,
		MIME:     fileMIME,
		Ext:      ext,
		IsDir:    f.IsDir(),
		Modified: f.ModTime(),
	}, nil
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

func (l *LocalStorageService) Search(path string, search string) ([]*model.File, error) {
	fullPath := filepath.Join(config.CFG.Storage.Path, path)
	rawFiles, err := readDir(fullPath)
	if err != nil {
		return nil, err
	}
	var Files []*model.File
	for _, f := range rawFiles {
		var fileType string
		var fileMIME string
		var ext string
		if !config.CFG.Storage.ShowHidden && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		if !f.IsDir() {
			fullFilePath := filepath.Join(fullPath, f.Name())
			fileType, fileMIME, err = util.FileType(fullFilePath)
			if err != nil {
				return nil, err
			}
			ext = util.FileExtension(f.Name())
		}
		if strings.Contains(strings.ToLower(f.Name()), strings.ToLower(search)) {
			Files = append(Files, &model.File{
				Name:     f.Name(),
				Path:     filepath.Join(path, f.Name()),
				Size:     f.Size(),
				Type:     fileType,
				MIME:     fileMIME,
				Ext:      ext,
				IsDir:    f.IsDir(),
				Modified: f.ModTime(),
			})
		}
	}
	return Files, nil
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
