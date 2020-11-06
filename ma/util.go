package ma

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/cliod/wx-go/common/util"
	"strings"
)

func CheckSignature(token, timestamp, nonce, signature string) bool {
	arr := []string{token, timestamp, nonce}
	gen, err := util.Gen(arr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return gen == signature
}

func GetUserInfo(sessionKey, encryptedData, ivStr string) (*UserInfo, error) {
	var usrInfo UserInfo
	err := util.Decrypt(&usrInfo, sessionKey, encryptedData, ivStr)
	return &usrInfo, err
}

func GetPhoneNoInfo(sessionKey, encryptedData, ivStr string) (*PhoneNumberInfo, error) {
	var info PhoneNumberInfo
	err := util.Decrypt(&info, sessionKey, encryptedData, ivStr)
	return &info, err
}

func CheckUserInfo(sessionKey, rawData, signature string) bool {
	sum := sha1.Sum([]byte(rawData + sessionKey))
	return strings.ToLower(hex.EncodeToString(sum[:])) == signature
}

func CheckAndGetUserInfo(sessionKey, rawData, encryptedData, signature, ivStr string) (*UserInfo, error) {
	var usrInfo UserInfo
	err := util.DecryptInfo(&usrInfo, sessionKey, rawData, encryptedData, signature, ivStr)
	return &usrInfo, err
}
