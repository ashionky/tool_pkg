/**
 * @Author pibing
 * @create 2022/1/25 9:39 AM
 */

package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"hash/crc32"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// UUID 产生uuid值
func UUID() string {
	uid := uuid.NewV1()
	return uid.String()
}

// Md5 该方法返回32位md5值
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// Md5With16 该方法返回md5加密值的16位值
func Md5With16(str string) string {
	return Md5(str)[8:24]
}

// Md5ToUpper 该方法返回md5加密值大写
func Md5ToUpper(str string) string {
	return strings.ToUpper(Md5(str))
}

// GetHashCodeByNum  计算字符串的ChecksumIEEE值与某个数的余数
func HashCodeByNum(str string, num int) int {
	v := crc32.ChecksumIEEE([]byte(str))
	if v < 0 {
		v = -v
	}
	return int(v) % num
}

// 带key的加密、解密
const HEX string = "0123456789ABCDEF"

func Encrypt(plainText string, key string) (string, error) {
	if len(key) != 32 {
		return "", errors.New("key length is not right")
	}
	keyBytes := []byte(key)
	plainBytes := []byte(plainText)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	ecb := cipher.NewCBCEncrypter(block, []byte(HEX))

	plainBytes = PKCS5Padding(plainBytes, block.BlockSize())
	crypted := make([]byte, len(plainBytes))
	ecb.CryptBlocks(crypted, plainBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil

}
func Decrypt(encrypted string, key string) (string, error) {
	if len(key) != 32 {
		return "", errors.New("key length is not right")
	}
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	encryptedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	ecb := cipher.NewCBCDecrypter(block, []byte(HEX))
	decryptedBytes := make([]byte, len(encryptedBytes))
	ecb.CryptBlocks(decryptedBytes, encryptedBytes)

	decryptedBytes = PKCS5Trimming(decryptedBytes)
	return string(decryptedBytes[:]), nil
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
