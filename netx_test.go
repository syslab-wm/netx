package netx

import (
	"testing"
)

func TestHasPort(t *testing.T) {
	data := []struct {
		addr     string // also the name of the test
		expected bool
	}{
		{"1.2.3.4:8080", true},
		{"1.2.3.4", false},
		{"1.2.3.4:", false},
	}

	for _, d := range data {
		t.Run(d.addr, func(t *testing.T) {
			got := HasPort(d.addr)
			if got != d.expected {
				t.Errorf("expected %t, got %t", d.expected, got)
			}
		})
	}
}

func TestTryJoinHostPort(t *testing.T) {
	data := []struct {
		server   string // also the name of the test
		port     string
		expected string
	}{
		{"1.2.3.4:8080", "8443", "1.2.3.4:8080"},
		{"1.2.3.4", "8443", "1.2.3.4:8443"},
		{"1.2.3.4:", "8443", "1.2.3.4:8443"},
	}

	for _, d := range data {
		t.Run(d.server, func(t *testing.T) {
			got := TryJoinHostPort(d.server, d.port)
			if got != d.expected {
				t.Errorf("expected %s, got %s", d.expected, got)
			}
		})
	}
}

func TestIsIPv4(t *testing.T) {
	data := []struct {
		s        string // also the name of the test
		expected bool
	}{
		{"1.2.3.4", true},
		{"127.0.0.1", true},
		{"::1", false},
		{"::ffff:192.0.2.128", false},
		{"2001:db8::1", false},
	}

	for _, d := range data {
		t.Run(d.s, func(t *testing.T) {
			got := IsIPv4(d.s)
			if got != d.expected {
				t.Errorf("expected %t, got %t", d.expected, got)
			}
		})
	}
}

func TestIsIPv6(t *testing.T) {
	data := []struct {
		s        string // also the name of the test
		expected bool
	}{
		{"1.2.3.4", false},
		{"127.0.0.1", false},
		{"::1", true},
		{"::ffff:192.0.2.128", true},
		{"2001:db8::1", true},
	}

	for _, d := range data {
		t.Run(d.s, func(t *testing.T) {
			got := IsIPv6(d.s)
			if got != d.expected {
				t.Errorf("expected %t, got %t", d.expected, got)
			}
		})
	}
}
