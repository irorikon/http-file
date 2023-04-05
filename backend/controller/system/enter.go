/*
 * @Author: iRorikon
 * @Date: 2023-04-04 17:06:09
 * @FilePath: \http-file\backend\controller\system\enter.go
 */
package system

import "github.com/irorikon/http-file/service"

type ControllerGroup struct {
	SystemController
}

var (
	systemService = service.ServiceGroupAPP.SystemServiceGroup.SystemService
)
