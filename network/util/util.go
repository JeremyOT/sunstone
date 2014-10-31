package util

import (
	"fmt"
	"net"
)

func TunnelName(prefix *net.IPNet) string {
	return fmt.Sprintf("bt%02x%02x%02x", prefix.IP[len(prefix.IP)-4], prefix.IP[len(prefix.IP)-3], prefix.IP[len(prefix.IP)-2])
}
