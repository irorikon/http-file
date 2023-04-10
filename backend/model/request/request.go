/*
 * @Author: iRorikon
 * @Date: 2023-04-05 01:04:44
 * @FilePath: \http-file\backend\model\request\request.go
 */
package request

import "github.com/irorikon/http-file/config"

type FileReq struct {
	StorageType string `json:"storage_type"`
	Path        string `json:"path"`
	NewPath     string `json:"new_path"`
	Filename    string `json:"filename"`
	Keyword     string `json:"keyword"`
}

type DownloadReq struct {
	Password string `json:"password"`
}

type SearchReq struct {
	Keyword string `json:"keyword"`
	Dir     string `json:"dir"`
}

type VideoPreviewReq struct {
	FileID string `json:"file_id"`
}

type SystemReq struct {
	System config.SystemConfig `json:"config"`
}
