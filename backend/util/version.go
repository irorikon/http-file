/*
 * @Author: iRorikon
 * @Date: 2023-04-04 16:30:03
 * @FilePath: \backend\util\version.go
 */
package util

import (
	"fmt"
	"runtime"
)

var (
	version   string
	buildDate string
	commit    string
	goVersion string
	tpl       string = `http-file, version %s (branch: HEAD, revision: %s)
    build user:       irorikon@88.com
    build date:       %s
    go version:       %s
    platform:         %s
 `
)

func VersionTpl() string {
	return fmt.Sprintf(tpl, version, commit, buildDate, goVersion, runtime.GOOS+"/"+runtime.GOARCH)
}
