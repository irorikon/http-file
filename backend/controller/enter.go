/*
 * @Author: iRorikon
 * @Date: 2023-04-04 17:06:09
 * @FilePath: \http-file\backend\controller\enter.go
 */
package controller

import "github.com/irorikon/http-file/service"

var (
	systemService = service.ServiceGroupAPP.SystemService
	fileService   = service.ServiceGroupAPP.FileService
)

type EnterController struct {
	SystemController
	FileController
}

var ControllerAPP = new(EnterController)
