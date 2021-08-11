package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandSeq 获取长度n的随机字符串
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// IsAnyEmpty values是否有为空的
func IsAnyEmpty(values ...string) bool {
	for _, v := range values {
		if strings.Trim(v, " ") == "" {
			return true
		}
	}
	return false
}

// Md5 签名
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// TrimToUpper 字符串去掉两端空格并转大写
func TrimToUpper(str string) string {
	return strings.ToUpper(strings.Trim(str, " "))
}
