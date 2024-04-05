package netx

import (
	"net"
	"net/netip"
	"strings"

	"github.com/syslab-wm/mu"
)

// HostPort returns whether addr includes a port number (i.e.,
// is of the form HOST:PORT).  It handles a corner-case in [net.SplitHostPort]
// which returns an empty port for addresses of the form "1.2.3.4:".  For such
// addresses, HasPort returns false.
func HasPort(addr string) bool {
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		return false
	}

	// this deals with the corner-case of, e.g., "1.2.3.4:".  For
	// addresses of this form, net.SplitHostAddr does not return an
	// error, and returns an empty port string.
	if port == "" {
		return false
	}
	return true
}

// TryJoinHostPort checks whether the server string already has a port (i.e.,
// ends with ':<PORT>'.  If it does, then the function simply returns
// that string.  If it does not, it returns the server string with
// the port appended.
func TryJoinHostPort(server string, port string) string {
	if HasPort(server) {
		return server
	}

	sanitized := server
	if strings.HasSuffix(server, ":") && !strings.HasSuffix(server, "::") {
		sanitized = server[:len(server)-1]
	}

	return net.JoinHostPort(sanitized, port)
}

// IsIPv4 returns true iff the addr string represents an IPv4 address.
func IsIPv4(s string) bool {
	addr, err := netip.ParseAddr(s)
	return err == nil && addr.Is4()
}

// IsIPv6 returns true iff the addr string represents an IPv6 address.
func IsIPv6(s string) bool {
	addr, err := netip.ParseAddr(s)
	return err == nil && addr.Is6()
}

// AddrAsIp converts a [netip.Addr] to a [net.IP].
func AddrAsIP(addr netip.Addr) net.IP {
	ip := net.ParseIP(addr.String())
	if ip == nil {
		mu.Panicf("can't convert netip.Addr (%v) as a netIP", addr)
	}
	return ip
}

// IPAsAddr converts a [net.IP] to a [netip.Addr].
func IPAsAddr(ip net.IP) netip.Addr {
	addr, err := netip.ParseAddr(ip.String())
	if err != nil {
		mu.Panicf("can't convert net.IP (%v) to a netip.Addr", ip)
	}
	return addr
}
