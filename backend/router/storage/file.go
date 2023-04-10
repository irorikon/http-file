/*
 * @Author: iRorikon
 * @Date: 2023-04-07 22:56:25
 * @FilePath: \http-file\backend\router\storage\file.go
 */
package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/controller"
)

type FileRouter struct{}

func (fr *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file")
	fileController := controller.ControllerAPP.File.FileController
	{
		fileRouter.POST("list", fileController.GetFileList)
		fileRouter.POST("info", fileController.GetFileInfo)
		fileRouter.POST("mkdir", fileController.MakeDir)
		fileRouter.POST("copy", fileController.Cpoy)
		fileRouter.POST("move", fileController.Move)
		fileRouter.POST("remove", fileController.Remove)
		fileRouter.POST("search", fileController.Search)
		fileRouter.GET("d/*link", fileController.Download)
	}
}
