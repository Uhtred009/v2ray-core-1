// +build !confonly

package stats

//go:generate go run v2ray.com/core/common/errors/errorgen

import (
	
	"sync"
	"bytes"
    "net"
)

type IPStorager struct {
	access sync.RWMutex
	ips []net.IP
}

func (s *IPStorager) Add(ip net.IP) bool {
	s.access.Lock()
	defer s.access.Unlock()

	for _, _ip := range s.ips {
		if bytes.Equal(_ip, ip) {
			return false
		}
	}

	s.ips = append(s.ips, ip)

	return true
}

func (s *IPStorager) Empty() {
	s.access.Lock()
	defer s.access.Unlock()

	s.ips = s.ips[:0]
}

func (s *IPStorager) Remove(removeIP net.IP) bool {
	s.access.Lock()
	defer s.access.Unlock()

	for i, ip := range s.ips {
		if bytes.Equal(ip, removeIP) {
			s.ips = append(s.ips[:i], s.ips[i+1:]...)
			return true
		}
	}

	return false
}

func (s *IPStorager) All() []net.IP {
	s.access.RLock()
	defer s.access.RUnlock()

	newIPs := make([]net.IP, len(s.ips))
	copy(newIPs, s.ips)

	return newIPs
}
