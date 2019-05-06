# IPMI

[![Build Status](https://travis-ci.org/gebn/ipmi.svg?branch=master)](https://travis-ci.org/gebn/ipmi)
[![GoDoc](https://godoc.org/github.com/gebn/ipmi?status.svg)](https://godoc.org/github.com/gebn/ipmi)
[![Go Report Card](https://goreportcard.com/badge/github.com/gebn/ipmi)](https://goreportcard.com/report/github.com/gebn/ipmi)

This project aims to implement a subset of the [IPMI](https://www.intel.co.uk/content/www/uk/en/servers/ipmi/ipmi-home.html) [v2.0 Specification](https://www.intel.com/content/dam/www/public/us/en/documents/specification-updates/ipmi-intelligent-platform-mgt-interface-spec-2nd-gen-v2-0-spec-update.pdf) in pure Go.
While IPMI is effectively deprecated, [Redfish](https://www.dmtf.org/standards/redfish) does not yet have enough support to be a viable replacement.
