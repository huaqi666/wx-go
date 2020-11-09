package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"sort"
	"strings"
)

// 对加密数据包进行签名校验，确保数据的完整性。
func validateSignature(signature string, parts ...string) bool {
	return signature == createSignature(parts...)
}

// 校验用户数据数据
func validateUserInfo(signature, rawData, ssk string) bool {
	return validateSignature(signature, rawData, ssk)
}

// 拼凑签名
func createSignature(parts ...string) string {
	sort.Strings(parts)
	raw := sha1.Sum([]byte(strings.Join(parts, "")))

	return hex.EncodeToString(raw[:])
}

// CBC 加密数据
func cbcEncrypt(key, plaintext, iv []byte) ([]byte, error) {
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv = iv[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// CBC解密数据
func cbcDecrypt(key, ciphertext, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := aes.BlockSize
	iv = iv[:size]
	// ciphertext = ciphertext[size:]

	if len(ciphertext) < size {
		return nil, errors.New("ciphertext too short")
	}

	if len(ciphertext)%size != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return pkcs7decode(ciphertext), nil
}

// 解密用户数据
func decryptData(ssk, ciphertext, iv string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return nil, err
	}

	cipherText, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	rawIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	return cbcDecrypt(key, cipherText, rawIV)
}
