// +build linux

package native

import (
	"net"
	"syscall"

	"github.com/JeremyOT/sunstone/network/netlink"
	"github.com/JeremyOT/sunstone/network/util"
)

func (m *nativeManager) RemoveBridge(bridgeName string) (err error) {
	err = netlink.DeleteBridge(bridgeName)
	return
}

func (m *nativeManager) CreateBridge(bridgeName string, ipNet *net.IPNet) (err error) {
	err = netlink.CreateBridge(bridgeName, true)
	if err != nil {
		return
	}
	iface, err := net.InterfaceByName(bridgeName)
	if err != nil {
		return
	}
	err = netlink.NetworkLinkAddIp(iface, ipNet.IP, ipNet)
	return
}

func (m *nativeManager) CreateTunnel(prefix *net.IPNet, remoteAddr net.IP, linkAddr net.IP, remoteLinkAddr net.IP) (err error) {
	tunnelName := util.TunnelName(prefix)
	err = netlink.CreateTunnel(tunnelName, remoteAddr, nil, syscall.IPPROTO_GRE, 255)
	if err != nil {
		return
	}
	iface, err := net.InterfaceByName(tunnelName)
	if err != nil {
		return
	}
	err = netlink.NetworkLinkAddIpPeer(iface, linkAddr, remoteLinkAddr, prefix)
	if err != nil {
		return
	}
	err = netlink.NetworkLinkUp(iface)
	if err != nil {
		return
	}
	_, safePrefix, err := net.ParseCIDR(prefix.String())
	if err != nil {
		return err
	}
	err = netlink.AddRoute(safePrefix.String(), "", "", tunnelName)
	if err != nil {
		return
	}
	return
}

func (m *nativeManager) RemoveTunnel(prefix *net.IPNet) (err error) {
	name := util.TunnelName(prefix)
	err = m.RemoveTunnelName(name)
	return
}

func (m *nativeManager) RemoveTunnelName(name string) (err error) {
	err = netlink.DeleteTunnel(name, 0)
	return
}
