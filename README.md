# netaddr
A Go library for performing calculations on IPv4 and IPv6 subnets. There is also limited support for EUI addresses.


# Installation
It should be noted that while the repository name is "netaddr-go" the package name is simply "netaddr".

	go get github.com/dspinhirne/netaddr-go


# Usage

	package main

	import "fmt"
	import "github.com/dspinhirne/netaddr-go"

	func main() {
		net,_ := netaddr.ParseIPv4Net("192.168.1.0/24")
		fmt.Println(net)
	}


# Documentation
Available online [here](https://godoc.org/github.com/dspinhirne/netaddr-go).

# Note
This is a fork of [dspinhirne/netaddr-go](https://github.com/dspinhirne/netaddr-go) for
personal experimentation. Please use the upstream project. This is not a stable
go package.
