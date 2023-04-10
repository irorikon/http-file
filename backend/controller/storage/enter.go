/*
 * @Author: iRorikon
 * @Date: 2023-04-05 22:34:33
 * @FilePath: \http-file\backend\controller\storage\enter.go
 */
package storage

import "github.com/irorikon/http-file/service"

type ControllerGroup struct {
	StorageController
	FileController
}

var (
	storageService      = service.ServiceGroupAPP.StorageServiceGroup.StorageService
	localStorageService = service.ServiceGroupAPP.StorageServiceGroup.LocalStorageService
)
