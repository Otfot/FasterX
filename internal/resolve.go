package internal

import (
	"net"
	"sync"

	"github.com/otfot/fasterx/pkg"
)

func GetRecord(src pkg.Source, multiDNS []string) *pkg.Record {
	ipSet := GetResolveIPSet(src, multiDNS)
	result := GetBestResult(ipSet, src)

	if result.Duration == pkg.MaxDuration {
		return nil
	}

	return &pkg.Record{
		Source: src,
		Result: result,
	}
}

func GetBestResult(ipSet map[string]net.IP, src pkg.Source) pkg.Result {
	var (
		task   = make(chan net.IP, 4)
		g      sync.WaitGroup
		worker = 5
		mutex  = new(sync.Mutex)
		best   = pkg.Result{
			Duration: pkg.MaxDuration,
		}
	)

	for i := 0; i < worker; i++ {
		g.Add(1)

		go func() {
			defer g.Done()
			for ip := range task {
				duration := pkg.ReqDuration(ip, src)
				mutex.Lock()
				if best.Duration > duration {
					best.Duration = duration
					best.IP = ip
				}
				mutex.Unlock()
			}
		}()
	}

	for _, ip := range ipSet {
		task <- ip
	}
	close(task)
	g.Wait()

	return best
}

func GetResolveIPSet(r pkg.Resolve, multiDNS []string) map[string]net.IP {
	var (
		ipSet  = make(map[string]net.IP)
		g      sync.WaitGroup
		tunnel = make(chan net.IP, 4)
	)

	go func() {
		for ip := range tunnel {
			ipSet[ip.String()] = ip
		}
		close(tunnel)
	}()

	for _, dnsServer := range multiDNS {
		g.Add(1)

		go func() {
			defer g.Done()
			ips := r.ResolveIPs(dnsServer)
			for _, ip := range ips {
				tunnel <- ip
			}
		}()
	}

	g.Wait()

	return ipSet
}
