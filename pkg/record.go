package pkg

import (
	"fmt"
	"net"
	"time"

	"github.com/miekg/dns"
)

type Resolve interface {
	ResolveIPs(dns string) []net.IP
}

type Source struct {
	Domain string
	URL    string
}

func (s Source) ResolveIPs(dnsServer string) []net.IP {
	// 创建一个 DNS 消息
	m := new(dns.Msg)

	m.SetQuestion(dns.Fqdn(s.Domain), dns.TypeA) // 查询A记录

	// 使用指定的 DNS 服务器发送查询请求
	client := new(dns.Client)
	response, _, err := client.Exchange(m, dnsServer+":53") // DNS 服务器默认端口为53
	if err != nil {
		return nil
	}

	var ips []net.IP
	// 解析响应中的 IP 地址
	for _, ans := range response.Answer {
		if aRecord, ok := ans.(*dns.A); ok {
			ips = append(ips, aRecord.A)
		}
	}
	return ips
}

type Result struct {
	IP       net.IP
	Duration time.Duration
}

func (rr Result) String() string {
	return fmt.Sprintf("IP: %s, Duration: %dms", rr.IP.String(), rr.Duration.Milliseconds())
}

type Record struct {
	Source Source
	Result Result
}

func (r *Record) Output() string {
	return r.Result.IP.String() + " " + r.Source.Domain
}
