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

// NewTLSClientByPkc12 通过pkc12证书创建支持双向证书认证的 http.Client.
func NewTLSClientByPkc12(data []byte, pw string) (client *http.Client, err error) {
	// 将pkcs12证书转成pem
	cert, err := pkc12ToPerm(data, pw)
	if err != nil {
		return
	}
	// tls配置
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	// 带证书的客户端
	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:     3 * time.Minute,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 10 * time.Minute,
			}).DialContext,
			TLSClientConfig:       config,
			DisableCompression:    true,
			MaxIdleConns:          100,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	return
}
