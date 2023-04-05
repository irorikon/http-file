/*
 * @Author: iRorikon
 * @Date: 2023-04-04 17:06:09
 * @FilePath: \http-file\backend\controller\enter.go
 */
package controller

import (
	"github.com/irorikon/http-file/controller/storage"
	"github.com/irorikon/http-file/controller/system"
	"github.com/irorikon/http-file/service"
)

var (
	systemService = service.ServiceGroupAPP.SystemServiceGroup
	fileService   = service.ServiceGroupAPP.StorageServiceGroup
)

type ControllerGroup struct {
	SystemController system.ControllerGroup
	FileController   storage.ControllerGroup
}

var ControllerAPP = new(ControllerGroup)
