/**
 * @Author pibing
 * @create 2022/1/25 11:24 AM
 */

package util

import (
	"crypto/rand"
	"math/big"
)

// 从str中生成指定位数length的随机字符串
func RandString(length int, str string) string {
	var chars = []byte(str)
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

//随机整数 min到max之间，包含min，不包含max，min和max不能相等
func RandIntByMinToMax(min, max int64) (int64, error) {
	dice, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		return dice.Int64(), err
	}
	return dice.Int64() + min, nil
}
