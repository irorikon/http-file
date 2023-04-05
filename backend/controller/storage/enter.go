package storage

import "github.com/irorikon/http-file/service"

type ControllerGroup struct {
}

var (
	localStorageService = service.ServiceGroupAPP.StorageServiceGroup.LocalStorageService
)
