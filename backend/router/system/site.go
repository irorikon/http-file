package system

import (
	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/controller"
)

type SiteRouter struct{}

func (s *SiteRouter) InitSiteRouter(Router *gin.RouterGroup) {
	siteRouter := Router.Group("site")
	siteController := controller.ControllerAPP.System.SiteController
	{
		siteRouter.GET("/get", siteController.GetSiteConfig)
		siteRouter.POST("/update", siteController.UpdateSiteConfig)
	}
}
