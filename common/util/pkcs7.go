package util

import (
	"bytes"
	"errors"
)

const pkcs7BlockSize = 32

var (
	ErrInvalidBlockSize    = errors.New("invalid block size")
	ErrInvalidPKCS7Data    = errors.New("invalid PKCS7 data")
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

// pkcs7encode 对需要加密的明文进行填充补位
// plaintext 需要进行填充补位操作的明文
// 返回补齐明文字符串
func pkcs7encode(plaintext []byte) []byte {
	//计算需要填充的位数
	pad := pkcs7BlockSize - len(plaintext)%pkcs7BlockSize
	if pad == 0 {
		pad = pkcs7BlockSize
	}

	//获得补位所用的字符
	text := bytes.Repeat([]byte{byte(pad)}, pad)

	return append(plaintext, text...)
}

// pkcs7decode 对解密后的明文进行补位删除
// plaintext 解密后的明文
// 返回删除填充补位后的明文和
func pkcs7decode(plaintext []byte) []byte {
	ln := len(plaintext)

	// 获取最后一个字符的 ASCII
	pad := int(plaintext[ln-1])
	if pad < 1 || pad > pkcs7BlockSize {
		pad = 0
	}

	return plaintext[:(ln - pad)]
}

// 解压
// pkcs7Unpacked returns slice of the original data without padding
func Pkcs7Unpacked(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
}
