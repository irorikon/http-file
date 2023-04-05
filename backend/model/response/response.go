/*
 * @Author: iRorikon
 * @Date: 2023-04-05 01:02:05
 * @FilePath: \http-file\backend\model\response\response.go
 */
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileListResponse struct {
	FileList []FileResponse `json:"filelist"`
}

type FileResponse struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64  `json:"size"`
	Dir  bool   `json:"dir"`
	MD5  string `json:"md5"`
}

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

const (
	ERROR   = -1
	SUCCESS = 0
)

func Result(code int, message string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func OK(c *gin.Context) {
	Result(SUCCESS, "success", nil, c)
}

func OKWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, message, nil, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(SUCCESS, "success", data, c)
}

func OKWithDetailed(data any, message string, c *gin.Context) {
	Result(SUCCESS, message, data, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "fail", nil, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, message, nil, c)
}

func FailWithDetailed(data any, message string, c *gin.Context) {
	Result(ERROR, message, data, c)
}
