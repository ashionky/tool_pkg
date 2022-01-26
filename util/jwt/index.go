/**
 * @Author pibing
 * @create 2020/12/27 12:14 PM
 */

package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

//把对象生成token
func MakeToken(obj map[string]interface{}, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims(obj))
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

//解析token为对象
func ParseToken(tokenStr, secret string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	finToken := token.Claims.(jwt.MapClaims)
	return finToken, nil
}
