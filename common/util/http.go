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

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return newTLSClient(config)
}

func newTLSClient(config *tls.Config) (*http.Client, error) {

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSClientConfig:       config,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{Transport: tr}
	//启用cookie
	//client.Jar, _ = cookiejar.New(nil)
	return client, nil
}
