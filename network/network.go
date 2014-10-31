package network

import (
	"net"

	"github.com/JeremyOT/sunstone/network/cli"
)

// Manager is responsible for creating and configuring network bridges and tunnels.
type Manager interface {
	// RemoveBridge removes an existing bridge with the given name.
	RemoveBridge(bridgeName string) (err error)
	// CreateBridge creates a bridge with the given name assigns it the given IP and subnet.
	CreateBridge(bridgeName string, ipNet *net.IPNet) (err error)
	// CreateTunnel creates a tunnel to remoteAddr between the given link addresses and configures it to handle traffic to and from the given subnet.
	CreateTunnel(prefix *net.IPNet, remoteAddr net.IP, linkAddr net.IP, remoteLinkAddr net.IP) (err error)
	// RemoveTunnel removes a tunnel created by CreateTunnel.
	RemoveTunnel(prefix *net.IPNet) (err error)
	// RemoveTunnelName removes an existing tunnel with the given name.
	RemoveTunnelName(name string) (err error)
}

func Default() Manager {
	return cli.New()
}
