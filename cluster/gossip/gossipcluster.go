package gossip

import (
	"fmt"
	"net"

	"github.com/JeremyOT/sunstone/cluster"
	"github.com/hashicorp/memberlist"
)

type GossipCluster struct {
	left         chan *cluster.Node
	joined       chan *cluster.Node
	members      *memberlist.Memberlist
	name         string
	eventHandler cluster.EventHandler
}

func (c *GossipCluster) NotifyJoin(node *memberlist.Node) {
	if node.Name == c.name {
		return
	}
	ip, ipNet, _ := net.ParseCIDR(node.Name)
	ipNet.IP = ip
	n := &cluster.Node{RemoteIP: node.Addr, Net: ipNet}
	c.eventHandler.NodeJoined(n)
}

func (c *GossipCluster) NotifyLeave(node *memberlist.Node) {
	if node.Name == c.name {
		fmt.Println("local left")
		return
	}
	ip, ipNet, _ := net.ParseCIDR(node.Name)
	ipNet.IP = ip
	n := &cluster.Node{RemoteIP: node.Addr, Net: ipNet}
	c.eventHandler.NodeLeft(n)
}

func (c *GossipCluster) NotifyUpdate(node *memberlist.Node) {

}

func (c *GossipCluster) Members() (nodes []*cluster.Node) {
	members := c.members.Members()
	nodes = make([]*cluster.Node, 0, len(members))
	for _, node := range members {
		if node.Name == c.name {
			continue
		}
		ip, ipNet, _ := net.ParseCIDR(node.Name)
		ipNet.IP = ip
		n := &cluster.Node{RemoteIP: node.Addr, Net: ipNet}
		nodes = append(nodes, n)
	}
	return
}

func (c *GossipCluster) Shutdown() error {
	return c.members.Shutdown()
}

func (c *GossipCluster) Join(seeds []string) (int, error) {
	return c.members.Join(seeds)
}

func New(address net.IP, port int, subnet *net.IPNet, eventHandler cluster.EventHandler) (cluster *GossipCluster, err error) {
	config := memberlist.DefaultLANConfig()
	if address != nil {
		config.BindAddr = address.String()
		config.AdvertiseAddr = config.BindAddr
	}
	config.BindPort = port
	config.AdvertisePort = port
	cluster = &GossipCluster{
		name:         subnet.String(),
		eventHandler: eventHandler,
	}
	config.Name = cluster.name
	config.Events = cluster
	memberlist, err := memberlist.Create(config)
	if err != nil {
		return
	}
	cluster.members = memberlist
	if err != nil {
		return
	}
	return
}
