/*
 * @Author: iRorikon
 * @Date: 2023-04-04 17:06:22
 * @FilePath: \http-file\backend\service\enter.go
 */
package service

type ServiceGroup struct {
	SystemService
	FileService
}

var ServiceGroupAPP = new(ServiceGroup)
