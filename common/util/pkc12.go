package util

import (
	"crypto/tls"
	"encoding/pem"
	"golang.org/x/crypto/pkcs12"
)

// pkc12转pem
func pkc12ToPerm(data []byte, password string) (cert tls.Certificate, err error) {
	blocks, err := pkcs12.ToPEM(data, password)
	if err != nil {
		return
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	cert, err = tls.X509KeyPair(pemData, pemData)
	return
}
