package util

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// NewTLSClient 创建支持双向证书认证的 http.Client.
func NewTLSClient(certPath, keyPath string) (*http.Client, error) {

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	return newTLSClient(config)
}

// NewTLSClientByPkc12 通过pkc12证书创建支持双向证书认证的 http.Client.
func NewTLSClientByPkc12(data []byte, pw string) (*http.Client, error) {
	// 将pkcs12证书转成pem
	cert, err := pkc12ToPerm(data, pw)
	if err != nil {
		return nil, err
	}
	// tls配置
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	// 带证书的客户端
	return newTLSClient(config)
}

// 指定配置创建http.Client
func newTLSClient(config *tls.Config) (*http.Client, error) {

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSClientConfig:       config,
		MaxIdleConns:          100,
		DisableCompression:    true,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{Timeout: 30 * time.Second, Transport: tr}
	//启用cookie
	//client.Jar, _ = cookiejar.New(nil)
	return client, nil
}
