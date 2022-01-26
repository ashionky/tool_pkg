/**
 * @Author pibing
 * @create 2021/6/7 1:43 PM
 */

package validator

import (
	"fmt"
	"regexp"
)

//验证邮箱
func VerifyEmail(email string, length int) bool {

	if len(email) > length { //长度不合法
		fmt.Println("email length Illegal ")
		return false
	}
	pattern := "^[0-9A-Za-z][0-9A-Za-z\\.!#\\$%&'\\*\\+\\-/=\\?\\^_`\\{\\}\\|~]*@([0-9A-Za-z][0-9A-Za-z\\-_]*\\.)+[0-9A-Za-z][0-9A-Za-z\\-_]*$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//手机号验证
func VerifyPhone(p string) bool {
	rex := `^(1(([35][0-9])|[8][0-9]|[9][0-9]|[6][0-9]|[7][01356789]|[4][579]))\d{8}$`
	reg := regexp.MustCompile(rex)
	return reg.MatchString(p)
}

//验证md5加密后的密码
func VerifyPassword(password string, length int) bool {

	if len(password) != length {
		fmt.Println("password length Illegal ")
		return false
	}
	pattern := "^[A-Za-z0-9]+$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(password)
}
