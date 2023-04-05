/*
 * @Author: iRorikon
 * @Date: 2023-04-05 15:13:39
 * @FilePath: \http-file\backend\util\base64.go
 */
package util

import "golang.org/x/crypto/bcrypt"

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
