package etcd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/JeremyOT/sunstone/cluster"
	"github.com/coreos/go-etcd/etcd"
)

type clusterNode struct {
	Address string `json:"address"`
	Subnet  string `json:"subnet"`
}

func putToUrl(targetUrl, body, contentType string) error {
	if request, err := http.NewRequest("PUT", targetUrl, strings.NewReader(body)); err != nil {
		return err
	} else {
		request.Header.Set("Content-Type", contentType)
		if resp, err := http.DefaultClient.Do(request); err != nil {
			return err
		} else {
			defer resp.Body.Close()
		}
	}
	return nil
}

type EtcdCluster struct {
	quit          chan struct{}
	etcdAddresses []string
	nodes         map[string]*cluster.Node
	address       net.IP
	port          int
	subnet        *net.IPNet
	eventHandler  cluster.EventHandler
	pollInterval  time.Duration
	etcdClient    *etcd.Client
	localNodeData string
	directory     string
	ttl           time.Duration
}

func New(address net.IP, port int, subnet *net.IPNet, eventHandler cluster.EventHandler, directory string, pollInterval time.Duration, ttl time.Duration) *EtcdCluster {
	c := &EtcdCluster{
		address:      address,
		port:         port,
		subnet:       subnet,
		eventHandler: eventHandler,
		pollInterval: pollInterval,
		directory:    directory,
		ttl:          ttl,
		nodes:        map[string]*cluster.Node{},
	}
	localNode := clusterNode{Address: address.String(), Subnet: subnet.String()}
	localNodeData, _ := json.Marshal(localNode)
	c.localNodeData = string(localNodeData)
	return c
}

func (c *EtcdCluster) Members() []*cluster.Node {
	nodes := make([]*cluster.Node, 0, len(c.nodes))
	for _, node := range c.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}

func (c *EtcdCluster) Shutdown() (err error) {
	close(c.quit)
	return
}

func (c *EtcdCluster) updateNodes() (err error) {
	_, err = c.etcdClient.Set(path.Join(c.directory, fmt.Sprintf("%s-%d", c.address.String(), c.port)), c.localNodeData, uint64(c.ttl/time.Second))
	if err != nil {
		return
	}
	resp, err := c.etcdClient.Get(c.directory, true, true)
	if err != nil {
		return
	}
	nodes := make(map[string]*cluster.Node, len(resp.Node.Nodes))
	for _, etcdNode := range resp.Node.Nodes {
		if etcdNode.Value == c.localNodeData {
			continue
		}
		var node clusterNode
		err = json.Unmarshal([]byte(etcdNode.Value), &node)
		if err != nil {
			return
		}
		nodeIP := net.ParseIP(node.Address)
		subnetIP, nodeSubnet, err := net.ParseCIDR(node.Subnet)
		if err != nil {
			return err
		}
		nodeSubnet.IP = subnetIP
		nodes[node.Subnet] = &cluster.Node{RemoteIP: nodeIP, Net: nodeSubnet}
	}
	removed := []*cluster.Node{}
	added := []*cluster.Node{}
	for key, node := range c.nodes {
		if _, ok := nodes[key]; !ok {
			removed = append(removed, node)
			delete(c.nodes, key)
		}
	}
	for key, node := range nodes {
		if _, ok := c.nodes[key]; !ok {
			if node.Net.String() == c.subnet.String() {
				continue
			}
			added = append(added, node)
			c.nodes[key] = node
		}
	}
	for _, node := range removed {
		c.eventHandler.NodeLeft(node)
	}
	for _, node := range added {
		c.eventHandler.NodeJoined(node)
	}
	return
}

func (c *EtcdCluster) pollNodes() {
	timer := time.Tick(c.pollInterval)
	for {
		select {
		case <-timer:
			err := c.updateNodes()
			if err != nil {
				log.Println("[Error] Error updating nodes:", err)
			}
		case <-c.quit:
			return
		}
	}
}

func (c *EtcdCluster) Join(seeds []string) (nodeCount int, err error) {
	if c.quit != nil {
		return 0, errors.New("Already joined.")
	}
	addresses := make([]string, 0, len(seeds))
	for _, address := range seeds {
		if address[:4] == "http" {
			addresses = append(addresses, address)
		} else {
			addresses = append(addresses, "http://"+address)
		}
	}
	c.quit = make(chan struct{})
	c.etcdAddresses = addresses
	c.etcdClient = etcd.NewClient(c.etcdAddresses)
	err = c.updateNodes()
	if err != nil {
		return
	}
	nodeCount = len(c.Members())
	go c.pollNodes()
	return
}
