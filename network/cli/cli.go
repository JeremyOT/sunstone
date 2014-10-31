package cli

import (
	"net"
	"os/exec"

	"github.com/JeremyOT/sunstone/network/util"
)

func cmd(args ...string) (err error) {
	path, err := exec.LookPath(args[0])
	if err != nil {
		return
	}
	return (&exec.Cmd{Path: path, Args: args}).Run()
}

type CliManager struct{}

func New() *CliManager {
	return &CliManager{}
}

func (c *CliManager) RemoveBridge(bridgeName string) (err error) {
	err = cmd("ip", "link", "set", bridgeName, "down")
	if err != nil {
		return
	}
	err = cmd("brctl", "delbr", bridgeName)
	return
}

func (c *CliManager) CreateBridge(bridgeName string, ipNet *net.IPNet) (err error) {
	c.RemoveBridge(bridgeName)
	err = cmd("brctl", "addbr", bridgeName)
	if err != nil {
		return
	}
	err = cmd("ip", "addr", "add", ipNet.String(), "dev", bridgeName)
	if err != nil {
		return
	}
	err = cmd("ip", "link", "set", bridgeName, "up")
	if err != nil {
		return
	}
	return
}

func (c *CliManager) CreateTunnel(prefix *net.IPNet, remoteAddr net.IP, linkAddr net.IP, remoteLinkeAddr net.IP) (err error) {
	name := util.TunnelName(prefix)
	err = cmd("ip", "tunnel", "add", name, "mode", "gre", "remote", remoteAddr.String())
	if err != nil {
		return
	}
	err = cmd("ip", "addr", "add", linkAddr.String(), "peer", remoteLinkeAddr.String(), "dev", name)
	if err != nil {
		return
	}
	err = cmd("ip", "link", "set", name, "up")
	if err != nil {
		return
	}
	_, safePrefix, err := net.ParseCIDR(prefix.String())
	if err != nil {
		return
	}
	err = cmd("ip", "route", "add", safePrefix.String(), "dev", name, "protocol", "static")
	if err != nil {
		return
	}
	return
}

func (c *CliManager) RemoveTunnel(prefix *net.IPNet) (err error) {
	name := util.TunnelName(prefix)
	err = c.RemoveTunnelName(name)
	return
}

func (c *CliManager) RemoveTunnelName(name string) (err error) {
	err = cmd("ip", "tunnel", "del", name)
	return
}
