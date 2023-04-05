/*
 * @Author: iRorikon
 * @Date: 2023-04-05 15:13:46
 * @FilePath: \http-file\backend\util\md5.go
 */
package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Str(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
