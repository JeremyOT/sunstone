package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/JeremyOT/address/lookup"
	"github.com/JeremyOT/sunstone/network/cli"
	"github.com/JeremyOT/sunstone/network/dummy"
	"github.com/JeremyOT/sunstone/network/native"
	"github.com/JeremyOT/sunstone/sunstone"
)

var bridgeName = flag.String("b", "", "Auto configure the specified bridge. Will create a new bridge if one does not exist. e.g. '-b docker0'")
var port = flag.Int("p", 9942, "The port to bind to for gossip communication.")
var localSubnetString = flag.String("n", "", "The subnet to map to this node.")
var tunnelSubnetString = flag.String("t", "", "The subnet to use for creating tunnels to other nodes.")
var seeds = flag.String("s", "", "The seeds to connect to, if any. Required when using Etcd.")
var runsunstone = flag.Bool("sunstone", true, "Set to false to only run setup operations.")
var networkDriver = flag.String("driver", "linux", "The driver to use for networking.")
var etcdTtl = flag.Duration("ttl", 15*time.Second, "The time in seconds that an entry will live in Etcd without update. Should be at least a few times longer than -interval.")
var etcdPollInterval = flag.Duration("interval", 3*time.Second, "How often to poll Etcd in seconds.")
var etcdKey = flag.String("etcd", "", "The key to use to track the cluster in Etcd. If set, Etcd will be used to maintain the cluster instead of gossip.")
var controlAddress = flag.String("control", "127.0.0.1:9842", "The address to use for control commands. Sunstone will bind to this address when starting a service or use it when running commands.")
var command = flag.Bool("command", false, "Send commands only.")
var join = flag.String("join", "", "The nodes to send in a join command.")
var reset = flag.Bool("reset", false, "Send a reset command to the cluster.")

func safeDefaults() {
	if *localSubnetString == "" {
		localAddress, err := lookup.GetAddress(true)
		if err != nil {
			log.Panicln(err)
		}
		localIP := net.ParseIP(localAddress).To4()
		localSubnet := net.IPv4(localIP[0], localIP[1]+1, 0, 0).String() + "/16"
		localSubnetString = &localSubnet
	}
	if *tunnelSubnetString == "" {
		localAddress, err := lookup.GetAddress(true)
		if err != nil {
			log.Panicln(err)
		}
		localIP := net.ParseIP(localAddress)
		subnetIP, _, err := net.ParseCIDR(*localSubnetString)
		if err != nil {
			log.Panicln(err)
		}
		tunnelIP := net.IPv4(subnetIP[len(subnetIP)-4], subnetIP[len(subnetIP)-3]+1, 0, 0)
		if tunnelIP[len(tunnelIP)-3] == localIP[len(localIP)-3] {
			if tunnelIP[len(tunnelIP)-3] > 149 {
				tunnelIP[len(tunnelIP)-3] -= 2
			} else {
				tunnelIP[len(tunnelIP)-3] += 1
			}
		}
		tunnelSubnet := tunnelIP.String() + "/16"
		tunnelSubnetString = &tunnelSubnet
	}
}

func monitorSignal(c <-chan os.Signal, m *sunstone.Sunstone) {
	for sig := range c {
		log.Println("Received signal:", sig)
		m.Shutdown()
		os.Exit(0)
	}
}

func main() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	flag.Usage = func() {
		fmt.Println("Usage")
		fmt.Println("  Start a Sunstone service: sunstone -b docker0")
		fmt.Println("  Control a Sunstone service: sunstone -command -join <seeds>")
		fmt.Println()
		fmt.Println("  See https://github.com/JeremyOT/sunstone for more information.")
		fmt.Println()
		fmt.Println("Options")
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()

	if *command {
		if *join != "" {
			err := sunstone.CommandJoin(*controlAddress, strings.Split(*join, ","))
			if err != nil {
				log.Panicln(err)
			}
		} else if *reset {
			err := sunstone.CommandReset(*controlAddress)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			log.Println("No command specified")
		}
		return
	}
	safeDefaults()
	ip, localSubnet, _ := net.ParseCIDR(*localSubnetString)
	localSubnet.IP = ip
	if *bridgeName != "" {
		if bridgeSubnet, err := sunstone.LocalSubnet(localSubnet); err != nil {
			log.Panicln(err)
		} else {
			if err := cli.New().CreateBridge(*bridgeName, bridgeSubnet); err != nil {
				log.Panicln(err)
			}
			log.Printf("Configured bridge %s for %s\n", *bridgeName, bridgeSubnet)
		}
	}
	ip, tunnelSubnet, _ := net.ParseCIDR(*tunnelSubnetString)
	tunnelSubnet.IP = ip
	localAddress, err := lookup.GetAddress(true)
	if err != nil {
		log.Panicln(err)
	}
	if !*runsunstone {
		return
	}
	var m *sunstone.Sunstone
	if *etcdKey != "" {
		m, err = sunstone.NewEtcd(*port, *etcdKey, *etcdPollInterval, *etcdTtl, localSubnet, tunnelSubnet)
		if err != nil {
			log.Panicln(err)
		}
	} else {
		m, err = sunstone.New(*port, localSubnet, tunnelSubnet)
		if err != nil {
			log.Panicln(err)
		}
	}
	if *networkDriver == "cli" {
		m.SetNetworkManager(cli.New())
	} else if *networkDriver == "dummy" {
		m.SetNetworkManager(dummy.New())
	} else {
		m.SetNetworkManager(native.New())
	}
	go m.Monitor(time.Minute, make(chan struct{}))
	go monitorSignal(sigChan, m)
	if *seeds != "" {
		nodes, err := m.Join(strings.Split(*seeds, ","))
		if err != nil {
			log.Panicln(err)
		}
		log.Println("Joined", nodes, "nodes")
	}
	fmt.Printf("Listening on %s:%d\n", localAddress, *port)
	log.Fatal(m.ListenForControl(*controlAddress))
}
