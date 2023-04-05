/*
 * @Author: iRorikon
 * @Date: 2023-04-05 00:58:46
 * @FilePath: \http-file\backend\controller\file.go
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/model/request"
	"github.com/irorikon/http-file/model/response"
)

type FileController struct {
}

func (fc *FileController) GetFileList(c *gin.Context) {
	var pathReq request.PathReq
	if err := c.ShouldBind(&pathReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pathReq.Password != config.CFG.Server.Password {
		response.FailWithMessage("密码错误", c)
		return
	}
	files, err := fileService.GetFileList(pathReq.Path)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithDetailed(files, "获取文件列表成功", c)
}
