package sunstone

import (
	"log"
	"net"
	"time"

	"github.com/JeremyOT/address/lookup"
	"github.com/JeremyOT/sunstone/cluster"
	"github.com/JeremyOT/sunstone/cluster/etcd"
	"github.com/JeremyOT/sunstone/cluster/gossip"
	"github.com/JeremyOT/sunstone/network"
	"github.com/JeremyOT/sunstone/network/util"
)

type Sunstone struct {
	cluster      cluster.Cluster
	localSubnet  *net.IPNet
	tunnelSubnet *net.IPNet
	netManager   network.Manager
}

func (m *Sunstone) SetNetworkManager(manager network.Manager) {
	m.netManager = manager
}

func (m *Sunstone) NodeJoined(node *cluster.Node) {
	log.Println("Node joined:", node)
	name := util.TunnelName(node.Net)
	tunnelIP, remoteTunnelIP := m.createTunnelIP(node)
	log.Println("Tunnel address:", tunnelIP)
	err := m.netManager.CreateTunnel(node.Net, node.RemoteIP, tunnelIP, remoteTunnelIP)
	if err != nil {
		log.Printf("[Error] Error creating tunnel %s: %s\n", name, err)
	} else {
		log.Println("Created tunnel", name)
	}
}

func (m *Sunstone) NodeLeft(node *cluster.Node) {
	log.Println("Node left:", node)
	name := util.TunnelName(node.Net)
	err := m.netManager.RemoveTunnelName(name)
	if err != nil {
		log.Printf("[Error] Error removing tunnel%s: %s\n", name, err)
	} else {
		log.Println("Removed tunnel", name)
	}
}

func LocalSubnet(subnet16 *net.IPNet) (localSubnet *net.IPNet, err error) {
	localAddress, err := lookup.GetAddress(true)
	if err != nil {
		return
	}
	localIP := net.ParseIP(localAddress)
	localSubnet = &net.IPNet{
		IP:   net.IPv4(subnet16.IP[len(subnet16.IP)-4], subnet16.IP[len(subnet16.IP)-3], localIP[len(localIP)-1], 1),
		Mask: net.IPv4Mask(255, 255, 255, 0),
	}
	return
}

func NewEtcd(port int, key string, pollInterval time.Duration, ttl time.Duration, subnet16 *net.IPNet, tunnelSubnet *net.IPNet) (sunstone *Sunstone, err error) {
	localSubnet, err := LocalSubnet(subnet16)
	if err != nil {
		return
	}
	sunstone = &Sunstone{localSubnet: localSubnet, tunnelSubnet: tunnelSubnet}
	localAddress, err := lookup.GetAddress(true)
	if err != nil {
		return
	}
	c := etcd.New(net.ParseIP(localAddress), port, localSubnet, sunstone, key, pollInterval, ttl)
	sunstone.cluster = c
	sunstone.netManager = network.Default()
	return
}

func New(port int, subnet16 *net.IPNet, tunnelSubnet *net.IPNet) (sunstone *Sunstone, err error) {
	localSubnet, err := LocalSubnet(subnet16)
	if err != nil {
		return
	}
	sunstone = &Sunstone{localSubnet: localSubnet, tunnelSubnet: tunnelSubnet}
	c, err := gossip.New(nil, port, localSubnet, sunstone)
	if err != nil {
		return
	}
	sunstone.cluster = c
	sunstone.netManager = network.Default()
	return
}

func (m *Sunstone) Monitor(interval time.Duration, quit chan struct{}) {
	timer := time.Tick(interval)
	for {
		select {
		case <-quit:
			return
		case <-timer:
			activeTunnels, err := tunnelInterfaceNames()
			if err != nil {
				log.Printf("[Error] %s\n", err)
			}
			joined, left := diffTunnels(activeTunnels, m.cluster.Members())
			for _, name := range left {
				log.Println("Detected dead node:", name)
				err := m.netManager.RemoveTunnelName(name)
				if err != nil {
					log.Printf("[Error] Error removing tunnel%s: %s\n", name, err)
				} else {
					log.Println("Removed tunnel", name)
				}
			}
			for name, node := range joined {
				log.Println("Detected new node:", name)
				tunnelIP, remoteTunnelIP := m.createTunnelIP(node)
				err := m.netManager.CreateTunnel(node.Net, node.RemoteIP, tunnelIP, remoteTunnelIP)
				if err != nil {
					log.Printf("[Error] Error creating tunnel %s: %s\n", name, err)
				} else {
					log.Println("Created tunnel", name)
				}
			}
		}
	}
}

func (m *Sunstone) createTunnelIP(node *cluster.Node) (tunnelIP, remoteTunnelIP net.IP) {
	r := node.Net.IP[len(node.Net.IP)-2]
	l := m.localSubnet.IP[len(m.localSubnet.IP)-2]
	tunnelIP = net.IPv4(m.tunnelSubnet.IP[len(m.tunnelSubnet.IP)-4], m.tunnelSubnet.IP[len(m.tunnelSubnet.IP)-3], l, r)
	remoteTunnelIP = net.IPv4(m.tunnelSubnet.IP[len(m.tunnelSubnet.IP)-4], m.tunnelSubnet.IP[len(m.tunnelSubnet.IP)-3], r, l)
	return
}

func (m *Sunstone) Shutdown() (err error) {
	log.Println("sunstone shutting down")
	err = m.cluster.Shutdown()
	if err != nil {
		return err
	}
	activeTunnels, err := tunnelInterfaceNames()
	if err != nil {
		return err
	}
	_, left := diffTunnels(activeTunnels, []*cluster.Node{})
	for _, name := range left {
		err := m.netManager.RemoveTunnelName(name)
		if err != nil {
			log.Printf("[Error] Error removing tunnel%s: %s\n", name, err)
		} else {
			log.Println("Removed tunnel", name)
		}
	}
	return
}

func (m *Sunstone) Join(seeds []string) (int, error) {
	return m.cluster.Join(seeds)
}

func diffTunnels(activeTunnels []string, nodes []*cluster.Node) (joined map[string]*cluster.Node, left []string) {
	joined = map[string]*cluster.Node{}
	left = []string{}
	tunnels := map[string]bool{}
	for _, t := range activeTunnels {
		tunnels[t] = true
	}
	members := map[string]*cluster.Node{}
	for _, n := range nodes {
		members[util.TunnelName(n.Net)] = n
	}
	for n, _ := range tunnels {
		if _, ok := members[n]; !ok {
			left = append(left, n)
		}
	}
	for n, node := range members {
		if _, ok := tunnels[n]; !ok {
			joined[n] = node
		}
	}
	return
}

func tunnelInterfaceNames() (names []string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	names = make([]string, 0, len(interfaces))
	for _, iface := range interfaces {
		if len(iface.Name) == 8 && iface.Name[0:2] == "bt" {
			names = append(names, iface.Name)
		}
	}
	return
}
