// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package reverseaddr provides a Reverse function that returns the
// reversed hostname of an IP address suitable for DNSBL record
// lookup or an error if it fails to parse the IP address.
//
// The function is correct for both IPv4 and IPv6.
//
// The code is taken from Go standard lib net package private
// reverseaddr. The only changes are to rename (export) this
// function from reverseaddr to Reverse and to not append the
// ARPA PTR zone suffices.
package reverseaddr

import (
	"bytes"
	"net"
)

const hexDigit = "0123456789abcdef"

// Reverse returns the reversed hostname of the IP address addr suitable
// for DNSBL record lookup or an error if it fails to parse the IP address.
//
// The function is correct for both IPv4 and IPv6.
func Reverse(addr string) (reverse string, err error) {
	ip := net.ParseIP(addr)
	if ip == nil {
		return "", &net.DNSError{Err: "unrecognized address", Name: addr}
	}
	if ip.To4() != nil {
		return uitoa(uint(ip[15])) + "." + uitoa(uint(ip[14])) + "." + uitoa(uint(ip[13])) + "." + uitoa(uint(ip[12])), nil
	}
	// Must be IPv6
	buf := make([]byte, 0, len(ip)*4)
	// Add it, in reverse, to the buffer
	for i := len(ip) - 1; i >= 0; i-- {
		v := ip[i]
		buf = append(buf, hexDigit[v&0xF])
		buf = append(buf, '.')
		buf = append(buf, hexDigit[v>>4])
		buf = append(buf, '.')
	}
	// buf has a trailing "."
	buf = bytes.TrimRight(buf, ".")
	return string(buf), nil
}

// Convert unsigned integer to decimal string.
func uitoa(val uint) string {
	if val == 0 { // avoid string allocation
		return "0"
	}
	var buf [20]byte // big enough for 64bit value base 10
	i := len(buf) - 1
	for val >= 10 {
		q := val / 10
		buf[i] = byte('0' + val - q*10)
		i--
		val = q
	}
	// val < 10
	buf[i] = byte('0' + val)
	return string(buf[i:])
}
