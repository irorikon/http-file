/*
 * @Author: iRorikon
 * @Date: 2023-04-05 21:31:31
 * @FilePath: \http-file\backend\model\model.go
 */
package model

import "time"

type File struct {
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	Size     int64     `json:"size"`
	Type     string    `json:"type"`
	MIME     string    `json:"mime"`
	Ext      string    `json:"ext"`
	Category string    `json:"category"`
	Hash     string    `json:"hash"`
	IsDir    bool      `json:"is_dir"`
	Modified time.Time `json:"modified"`
	URL      string    `json:"url"`
}
