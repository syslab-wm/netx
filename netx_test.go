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

func TestTryAddPort(t *testing.T) {
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
			got := TryAddPort(d.server, d.port)
			if got != d.expected {
				t.Errorf("expected %s, got %s", d.expected, got)
			}
		})
	}
}
