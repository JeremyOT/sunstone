package dummy

import (
	"net"

	"github.com/JeremyOT/sunstone/network/util"
)

type dummyManager struct {
	bridges map[string]bool
	tunnels map[string]bool
}

func New() *dummyManager {
	return &dummyManager{
		bridges: map[string]bool{},
		tunnels: map[string]bool{},
	}
}

func (m *dummyManager) RemoveBridge(bridgeName string) (err error) {
	delete(m.bridges, bridgeName)
	return
}

func (m *dummyManager) CreateBridge(bridgeName string, ipNet *net.IPNet) (err error) {
	m.bridges[bridgeName] = true
	return
}

func (m *dummyManager) CreateTunnel(prefix *net.IPNet, remoteAddr net.IP, linkAddr net.IP, remoteLinkeAddr net.IP) (err error) {
	m.bridges[util.TunnelName(prefix)] = true
	return
}

func (m *dummyManager) RemoveTunnel(prefix *net.IPNet) (err error) {
	return m.RemoveTunnelName(util.TunnelName(prefix))
}

func (m *dummyManager) RemoveTunnelName(name string) (err error) {
	delete(m.tunnels, name)
	return
}

func (m *dummyManager) Tunnels() []string {
	tunnels := make([]string, 0, len(m.tunnels))
	for tunnel := range m.tunnels {
		tunnels = append(tunnels, tunnel)
	}
	return tunnels
}

func (m *dummyManager) Bridges() []string {
	bridges := make([]string, 0, len(m.bridges))
	for bridge := range m.bridges {
		bridges = append(bridges, bridge)
	}
	return bridges
}
