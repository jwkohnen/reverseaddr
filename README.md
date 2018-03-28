# ReverseAddr -- reverse an IP representation for DNSBL lookups

Package reverseaddr provides a Reverse function that returns the
reversed hostname of an IP address suitable for DNSBL record
lookup or an error if it fails to parse the IP address.

The function is correct for both IPv4 and IPv6.

The code is taken from Go standard lib net package private
reverseaddr. The only changes are to rename (export) this
function from reverseaddr to Reverse and to not append the
ARPA PTR zone suffices.

## License

Copyright 2009 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
