package service

import (
	"mime/multipart"
	"os"

	"github.com/irorikon/http-file/model/response"
	"github.com/irorikon/http-file/util"
)

type FileService struct{}

func (fs *FileService) GetFileList(filePath string) (response.FileListResponse, error) {
	var flr response.FileListResponse
	fl, err := os.ReadDir(filePath)
	if err != nil {
		return response.FileListResponse{}, err
	}
	for _, f := range fl {
		fileinfo, err := f.Info()
		if err != nil {
			return response.FileListResponse{}, err
		}
		var isDir bool
		var size int64
		if f.IsDir() {
			isDir = true
			size = 0
		} else {
			isDir = false
			size = fileinfo.Size()
		}
		flr.FileList = append(flr.FileList, response.FileResponse{
			Name: fileinfo.Name(),
			Size: size,
			Path: filePath,
			Dir:  isDir,
			MD5:  "",
		})
	}
	return flr, nil
}
func (fs *FileService) List(filePath string) (response.FileResponse, error) {
	fileinfo, err := os.Stat(filePath)
	if err != nil {
		return response.FileResponse{}, err
	}
	var isDir bool
	var size int64
	if fileinfo.IsDir() {
		isDir = true
		size = 0
	} else {
		isDir = false
		size = fileinfo.Size()
	}
	filebyte, err := os.ReadFile(filePath)
	if err != nil {
		return response.FileResponse{}, err

	}
	return response.FileResponse{
		Name: fileinfo.Name(),
		Size: size,
		Path: filePath,
		Dir:  isDir,
		MD5:  util.MD5Str(filebyte),
	}, nil
}

func (fs *FileService) Create(file *multipart.FileHeader) (string, error) {
	return "", nil
}

func (fs *FileService) Delete(filePath string) error {
	return nil
}

func (fs *FileService) Download(filePath string) (string, error) {
	return "", nil
}

func (fs *FileService) Upload(file *multipart.FileHeader) (string, error) {
	return "", nil
}
