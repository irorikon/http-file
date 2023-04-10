package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/controller"
)

type StorageRouter struct{}

func (sr *StorageRouter) InitStorageRoute(Router *gin.RouterGroup) {
	storageRouter := Router.Group("storage")
	storageController := controller.ControllerAPP.File.StorageController
	{
		storageRouter.GET("get", storageController.GetStorageConfig)
		storageRouter.POST("update", storageController.UpdateStorageConfig)
	}
}
