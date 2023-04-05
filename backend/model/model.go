/*
 * @Author: iRorikon
 * @Date: 2023-04-05 21:31:31
 * @FilePath: \http-file\backend\model\model.go
 */
package model

import "time"

type File struct {
	Name     string
	Path     string
	Size     int64
	Type     string
	Hash     string
	IsDir    bool
	Modified time.Time
}
