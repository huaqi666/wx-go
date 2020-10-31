package util

import (
	"encoding/json"
	"errors"
)

func Decrypt(v interface{}, sessionKey, encryptedData, ivStr string) error {
	raw, err := decryptData(sessionKey, encryptedData, ivStr)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, v)
}

// DecryptInfo 解密用户信息
//
// sessionKey 微信 session_key
// rawData 不包括敏感信息的原始数据字符串，用于计算签名。
// encryptedData 包括敏感数据在内的完整用户信息的加密数据
// signature 使用 sha1( rawData + session_key ) 得到字符串，用于校验用户信息
// iv 加密算法的初始向量
func DecryptInfo(v interface{}, sessionKey, rawData, encryptedData, signature, iv string) error {

	if ok := validateUserInfo(signature, rawData, sessionKey); !ok {
		return errors.New("failed to validate signature")
	}

	raw, err := decryptData(sessionKey, encryptedData, iv)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, v)
}
