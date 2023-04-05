/*
 * @Author: iRorikon
 * @Date: 2023-04-04 17:06:22
 * @FilePath: \http-file\backend\service\enter.go
 */
package service

import (
	"github.com/irorikon/http-file/service/storage"
	"github.com/irorikon/http-file/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	StorageServiceGroup storage.ServiceGroup
}

var ServiceGroupAPP = new(ServiceGroup)
