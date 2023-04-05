/*
 * @Author: iRorikon
 * @Date: 2023-04-05 01:04:44
 * @FilePath: \http-file\backend\model\request\request.go
 */
package request

import "github.com/irorikon/http-file/config"

type DownloadReq struct {
	Password string `json:"password"`
}

type GetReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
}

type OfficePreviewReq struct {
	FileID string `json:"file_id"`
}

type PathReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
}

type ProxyReq struct {
	Password string `json:"password"`
}

type SearchReq struct {
	Keyword string `json:"keyword"`
	Dir     string `json:"dir"`
}

type RebuildTreeReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
	Depth    int    `json:"depth"`
}

type VideoPreviewReq struct {
	FileID string `json:"file_id"`
}

type SystemReq struct {
	System config.SystemConfig `json:"config"`
}
