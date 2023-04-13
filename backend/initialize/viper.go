/*
 * @Author: iRorikon
 * @Date: 2023-04-05 02:16:22
 * @FilePath: \http-file\backend\initialize\viper.go
 */
package initialize

import (
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/irorikon/http-file/config"
	"github.com/irorikon/http-file/util"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var c string
	if len(path) == 0 {
		c = config.ConfigFile
		if c == "" {
			if configEnv := os.Getenv(config.ConfigEnv); configEnv == "" {
				c = config.DefaultConfigFile
				fmt.Printf("您正在使用 config 的默认值, config 的路径为: %v\n", config.DefaultConfigFile)
			} else {
				c = configEnv
				fmt.Printf("您正在使用 ConfigEnv 环境变量, config 的路径为: %v\n", c)
			}
		}
	} else {
		fmt.Printf("您正在使用命令行的 -c 参数传递的值, config 的路径为: %v\n", c)
	}

	if !util.FileExist(c) {
		fmt.Printf("file not exist: %s", c)
		// 文件不存在，创建文件
		err := viper.WriteConfigAs(c)
		if err != nil {
			panic(err)
		}
		err = util.WriteYamlToFile(c, config.DefaultConfig())
		if err != nil {
			panic(err)
		}
	}

	v := viper.New()
	v.SetConfigFile(c)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&config.CFG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&config.CFG); err != nil {
		fmt.Println(err)
	}
	config.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(config.CFG.JWT.Expire)),
	)
	return v
}
