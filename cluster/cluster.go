package cluster

import "net"

type Node struct {
	RemoteIP net.IP
	Net      *net.IPNet
}

type Cluster interface {
	Members() (nodes []*Node)
	Shutdown() error
	Join(seeds []string) (int, error)
}

type EventHandler interface {
	NodeJoined(node *Node)
	NodeLeft(node *Node)
}
