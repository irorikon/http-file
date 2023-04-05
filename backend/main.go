/*
 * @Author: iRorikon
 * @Date: 2023-04-04 16:48:14
 * @FilePath: \http-file\backend\main.go
 */
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/initialize"
	"github.com/irorikon/http-file/util"
)

func init() {
	flag.StringVar(&config.ConfigFile, "config-file", "", "Password.")
	flag.StringVar(&config.ConfigFile, "c", "", "Password.")
	flag.BoolVar(&config.Version, "v", false, "Show version and exit.")
	flag.BoolVar(&config.Version, "version", false, "Show version and exit.")
	flag.Usage = usage
}

var usageStr = `
Usage: http-file [option]

Optionnal parameters:
  -c, --config-file     Path to config file.
  -v, --version         Show version and exit.
  -h, --help            Show this help.
`

func usage() {
	fmt.Println(usageStr)
	os.Exit(0)
}

func main() {
	flag.Parse()
	if config.Version {
		fmt.Println(util.VersionTpl())
		os.Exit(0)
	}
	// 初始化Viper
	config.Viper = initialize.Viper()
	// 初始化zap日志库
	config.Log = initialize.Zap()
	// 运行服务器
	initialize.RunServer()
}
