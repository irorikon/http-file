//go:build windows
// +build windows

/*
 * @Author: iRorikon
 * @Date: 2023-04-05 03:00:55
 * @FilePath: \http-file\backend\initialize\server_win.go
 */

package initialize

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
