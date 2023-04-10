package system

import (
	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/controller"
)

type SystemRouter struct{}

func (s *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) {
	systemRouter := Router.Group("system")
	systemController := controller.ControllerAPP.System.SystemController
	{
		systemRouter.GET("/get", systemController.GetSystemConfig)
		systemRouter.POST("/update", systemController.UpdateSystemConfig)
		systemRouter.GET("/state", systemController.GetSystemState)
	}
}
