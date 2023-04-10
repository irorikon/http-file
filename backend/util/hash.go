/*
 * @Author: iRorikon
 * @Date: 2023-04-05 15:13:39
 * @FilePath: \http-file\backend\util\base64.go
 */
package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用 bcrypt 对字符串进行加密
func BcryptHash(str string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文字符串和哈希值
func BcryptCheck(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}

func MD5Str(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// Base64Encode 对字符串进行base64编码
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode 对字符串进行base64解码
func Base64Decode(str string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
