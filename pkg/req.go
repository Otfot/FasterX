package pkg

import (
	"context"
	"crypto/tls"
	"math"
	"net"
	"net/http"
	"time"
)

const (
	Timeout     = 5 * time.Second
	MaxDuration = math.MaxInt64
)

func newDirectClient(ip net.IP) *http.Client {
	directAddr := ip.String() + ":443"
	// 创建自定义 Transport，指定拨号器和跳过不受信任证书检查
	transport := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// 强制使用指定的 IP 地址，但保留原主机名（用于 TLS 证书验证和 Host 头）
			return (&net.Dialer{
				Timeout: Timeout,
			}).DialContext(ctx, network, directAddr)
		},
	}

	// 创建 HTTP 客户端，使用自定义传输
	client := &http.Client{
		Transport: transport,
		Timeout:   Timeout,
	}

	return client
}

func ReqDuration(ip net.IP, src Source) time.Duration {
	client := newDirectClient(ip)
	// 发起 HTTPS 请求
	// 创建请求
	req, err := http.NewRequest("GET", src.URL, nil)
	if err != nil {
		return MaxDuration
	}
	// 设置 Host 头，以确保请求中的主机名用于 TLS 验证
	req.Host = src.Domain
	t1 := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return MaxDuration
	}
	defer resp.Body.Close()
	t2 := time.Since(t1)
	return t2
}
