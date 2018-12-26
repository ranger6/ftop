package main

import (
	"fmt"
	"testing"
)

func TestFtop(t *testing.T) {
	var tests = []struct {
		in       string
		expected string
	}{
		{"www.hp.com", "com/hp/www"},
		{"www", "www"},
		{"", ""},
	}
	for _, test := range tests {
		got := ftop(test.in)
		if got != test.expected {
			t.Errorf("From %q, expected %q, but got %q", test.in, test.expected, got)
		}
	}
}

func ExampleFqdn() {
	fmt.Println(ftop("borax.fra.hp.com"))
	// Output:
	// com/hp/fra/borax
}
