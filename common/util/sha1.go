package util

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

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
