package pkg

// https://dns.icoa.cn/#china
const (
	BaiduDNS     = "180.76.76.76"
	AliDNS       = "223.5.5.5"
	II4DNS       = "114.114.114.114"
	TecentDNS    = "119.29.29.29"
	BytedanceDNS = "180.184.1.1"
	OneDNS       = "117.50.10.10"
)

// https://dns.icoa.cn/#world
const (
	DNSSBDNS        = "185.222.222.222"
	CloudflareDNS   = "1.1.1.1"
	GoogleDNS       = "8.8.8.8"
	Quad9DNS        = "9.9.9.9"
	OpenDNS         = "208.67.222.222"
	YandexDNS       = "77.88.8.8"
	AdGuardDNS      = "94.140.14.140"
	Level3DNS       = "4.2.2.1"
	FreenomWorldDNS = "80.80.80.80"
)

var InDNS = []string{
	BaiduDNS,
	AliDNS,
	II4DNS,
	TecentDNS,
	BytedanceDNS,
	OneDNS,
}

var OutDNS = []string{
	DNSSBDNS,
	CloudflareDNS,
	GoogleDNS,
	Quad9DNS,
	OpenDNS,
	YandexDNS,
	AdGuardDNS,
	Level3DNS,
	FreenomWorldDNS,
}
