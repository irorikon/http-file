/*
 * @Author: iRorikon
 * @Date: 2023-04-07 21:57:51
 * @FilePath: \http-file\backend\controller\storage\file.go
 */
package storage

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/model/request"
	"github.com/irorikon/http-file/model/response"
	"github.com/irorikon/http-file/util"
)

type FileController struct{}

// @function: GetFileList
// @description: 获取文件列表
func (fc *FileController) GetFileList(c *gin.Context) {
	var fileReq request.FileReq
	if err := c.ShouldBindJSON(&fileReq); err != nil {
		response.FailWithMessage("Parameter error", c)
		return
	}
	// 验证密码
	// TODO
	storageType := strings.ToLower(fileReq.StorageType)
	switch storageType {
	case "local":
		files, err := localStorageService.GetFileList(fileReq.Path)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OKWithDetailed(files, "success", c)
	default:
		response.FailWithMessage("The storage format is not currently supported", c)
	}
}

// @function: GetFileInfo
// @description: 获取文件信息
func (fc *FileController) GetFileInfo(c *gin.Context) {
	var fileReq request.FileReq
	if err := c.ShouldBindJSON(&fileReq); err != nil {
		response.FailWithMessage("Parameter error", c)
		return
	}
	// 验证密码
	// TODO
	storageType := strings.ToLower(fileReq.StorageType)
	switch storageType {
	case "local":
		file, err := localStorageService.GetFileInfo(fileReq.Path)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OKWithDetailed(file, "success", c)
	default:
		response.FailWithMessage("The storage format is not currently supported", c)
	}
}

// @function: MakeDir
// @description: 创建文件夹
func (fc *FileController) MakeDir(c *gin.Context) {
	var fileReq request.FileReq
	if err := c.ShouldBindJSON(&fileReq); err != nil {
		response.FailWithMessage("Parameter error", c)
		return
	}
	// 验证密码
	// TODO
	storageType := strings.ToLower(fileReq.StorageType)
	switch storageType {
	case "local":
		err := localStorageService.MakeDir(fileReq.Path)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OKWithMessage("success", c)
	default:
		response.FailWithMessage("The storage format is not currently supported", c)
	}
}

// @function: Move
// @description: 移动文件
func (fc *FileController) Move(c *gin.Context) {
	var fileReq request.FileReq
	if err := c.ShouldBindJSON(&fileReq); err != nil {
		response.FailWithMessage("Parameter error", c)
		return
	}
	// 验证密码
	// TODO
	storageType := strings.ToLower(fileReq.StorageType)
	switch storageType {
	case "local":
		err := localStorageService.Move(fileReq.NewPath, fileReq.Path)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OKWithMessage("success", c)
	default:
		response.FailWithMessage("The storage format is not currently supported", c)
	}
}

// @function: Cpoy
// @description: 复制文件
func (fc *FileController) Cpoy(c *gin.Context) {
	var fileReq request.FileReq
	if err := c.ShouldBindJSON(&fileReq); err != nil {
		response.FailWithMessage("Parameter error", c)
		return
	}
	// 验证密码
	// TODO
	storageType := strings.ToLower(fileReq.StorageType)
	switch storageType {
	case "local":
		err := localStorageService.Cpoy(fileReq.NewPath, fileReq.Path)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OKWithMessage("success", c)
	default:
		response.FailWithMessage("The storage format is not currently supported", c)
	}
}

// @function: Remove
// @description: 删除文件
func (fc *FileController) Remove(c *gin.Context) {
	var fileReq request.FileReq
	if err := c.ShouldBindJSON(&fileReq); err != nil {
		response.FailWithMessage("Parameter error", c)
		return
	}
	// 验证密码
	// TODO
	storageType := strings.ToLower(fileReq.StorageType)
	switch storageType {
	case "local":
		err := localStorageService.Remove(fileReq.Path)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OKWithMessage("success", c)
	default:
		response.FailWithMessage("The storage format is not currently supported", c)
	}
}

// @function: Download
// @description: 下载文件
func (fc *FileController) Download(c *gin.Context) {
	link, ok := c.Params.Get("link")
	if !ok {
		response.FailWithMessage("Parameter error", c)
		return
	}
	fullpath := filepath.Join(config.CFG.Storage.Path, link)
	fileinfo, err := os.Stat(fullpath)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if fileinfo.IsDir() {
		response.FailWithMessage("The file is a directory", c)
		return
	}
	c.File(fullpath)
}

// @function: DownloadEncode
// @description: 下载文件
func (fc *FileController) DownloadEncode(c *gin.Context) {
	link, ok := c.Params.Get("link")
	if !ok {
		response.FailWithMessage("Parameter error", c)
		return
	}
	filename, err := util.FileDownloadDecode(link[1:])
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	fullpath := filepath.Join(config.CFG.Storage.Path, filename)
	c.File(fullpath)
}

// @function: Search
// @description: 搜索
func (fc *FileController) Search(c *gin.Context) {
	var searchReq request.FileReq
	if err := c.ShouldBindJSON(&searchReq); err != nil {
		response.FailWithMessage("Parameter error", c)
		return
	}
	// 验证密码
	// TODO
	storageType := strings.ToLower(searchReq.StorageType)
	switch storageType {
	case "local":
		files, err := localStorageService.Search(searchReq.Path, searchReq.Keyword)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OKWithDetailed(files, "success", c)
	default:
		response.FailWithMessage("The storage format is not currently supported", c)
	}
}
