/*
 * @Author: iRorikon
 * @Date: 2023-04-05 22:46:36
 * @FilePath: \http-file\backend\middleware\load_tls.go
 */
package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/irorikon/http-file/config"
	"github.com/unrolled/secure"
)

func LoadTLS() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     fmt.Sprintf("%s:443", config.CFG.System.Address),
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			// 如果出现错误，请不要继续
			fmt.Println(err)
			return
		}
		// 继续往下处理
		c.Next()
	}
}
