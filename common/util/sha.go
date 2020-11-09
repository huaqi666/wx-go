package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// 组装url参数并签名
func GenWithAmple(arr []string) (string, error) {
	if IsAnyEmpty(arr...) {
		return "", fmt.Errorf("非法请求参数，有部分参数为空")
	}
	sort.Strings(arr)
	var str string
	for i, v := range arr {
		str += v
		if i != len(arr)-1 {
			str += "&"
		}
	}
	sum := sha1.Sum([]byte(str))
	return strings.ToLower(hex.EncodeToString(sum[:])), nil
}

// 组装url参数并签名
func Gen(arr []string) (string, error) {
	if IsAnyEmpty(arr...) {
		return "", fmt.Errorf("非法请求参数，有部分参数为空")
	}
	sort.Strings(arr)
	var str string
	for _, v := range arr {
		str += v
	}
	sum := sha1.Sum([]byte(str))
	return strings.ToLower(hex.EncodeToString(sum[:])), nil
}

// HmacSha256 签名
func HmacSha256(str, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(str))
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}
