package netlink

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <sys/ioctl.h>
#include <net/if.h>
#include <net/if_arp.h>
#include <linux/ip.h>
#include <linux/if_tunnel.h>
#include <errno.h>

struct ip_tunnel_parm makeParm(char *name, unsigned char protocol, unsigned char ttl, unsigned int remoteAddr, unsigned int localAddr, unsigned short fragOff) {
	struct ip_tunnel_parm p;
	memset(&p, 0, sizeof(p));
	p.iph.version = 4;
	p.iph.ihl = 5;
	p.iph.protocol = protocol;
	p.iph.frag_off = htons(0x4000);
	p.iph.ttl = ttl;
	p.iph.daddr = remoteAddr;
	p.iph.saddr = localAddr;
	strncpy(p.name, name, IFNAMSIZ);
	return p;
}

int tnl_add_ioctl(int cmd, const char *basedev, const char *name, struct ip_tunnel_parm p)
{
	struct ifreq ifr;
	int fd;
	int err;

	if (cmd == SIOCCHGTUNNEL && name[0])
		strncpy(ifr.ifr_name, name, IFNAMSIZ);
	else
		strncpy(ifr.ifr_name, basedev, IFNAMSIZ);
	ifr.ifr_ifru.ifru_data = (void*)&p;
	fd = socket(AF_INET, SOCK_DGRAM, 0);
	err = ioctl(fd, cmd, &ifr);
	close(fd);
	return errno;
}
*/
import "C"
import (
	"errors"
	"fmt"
	"net"
	"syscall"
	"unsafe"
)

const (
	SIOCDEVPRIVATE = 0x89F0
	SIOCGETTUNNEL  = SIOCDEVPRIVATE + 0
	SIOCADDTUNNEL  = SIOCDEVPRIVATE + 1
	SIOCDELTUNNEL  = SIOCDEVPRIVATE + 2
	SIOCCHGTUNNEL  = SIOCDEVPRIVATE + 3
)

type iphdr struct {
	ihl      byte
	version  byte
	tos      byte
	totLen   uint16
	id       uint16
	fragOff  uint16
	ttl      byte
	protocol byte
	check    uint16
	saddr    uint32
	daddr    uint32
}

func IPToUInt32(ip net.IP) uint32 {
	if ip == nil {
		return 0
	}
	return native.Uint32(ip.To4())
}

type ifreq struct {
	IfrnName [IFNAMSIZ]byte
	IfruData uintptr
}

func CreateTunnel(name string, remoteAddr, localAddr net.IP, protocol int, ttl byte) error {

	if len(name) >= IFNAMSIZ {
		return fmt.Errorf("Interface name %s too long", name)
	}

	s, err := getIfSocket()
	if err != nil {
		return err
	}
	defer syscall.Close(s)

	var baseDevice string
	switch protocol {
	case syscall.IPPROTO_GRE:
		baseDevice = "gre0"
	case syscall.IPPROTO_IPIP:
		baseDevice = "tunl0"
	case syscall.IPPROTO_IPV6:
		baseDevice = "sit0"
	default:
		return errors.New("Invalid tunnel mode.")
	}
	parms := C.makeParm(C.CString(name), C.uchar(protocol), C.uchar(ttl), C.uint(IPToUInt32(remoteAddr)), C.uint(IPToUInt32(localAddr)), 0)
	res := C.tnl_add_ioctl(SIOCADDTUNNEL, C.CString(baseDevice), C.CString(name), parms)
	if res != 0 {
		return errors.New(C.GoString(C.strerror(res)))
	}
	return nil
}

func DeleteTunnel(name string, protocol int) error {
	if len(name) >= IFNAMSIZ {
		return fmt.Errorf("Interface name %s too long", name)
	}

	s, err := getIfSocket()
	if err != nil {
		return err
	}
	defer syscall.Close(s)

	var baseDevice string
	switch protocol {
	case syscall.IPPROTO_GRE:
		baseDevice = "gre0"
	case syscall.IPPROTO_IPIP:
		baseDevice = "tunl0"
	case syscall.IPPROTO_IPV6:
		baseDevice = "sit0"
	default:
		baseDevice = name
	}

	ifr := ifreq{IfruData: uintptr(unsafe.Pointer(&[0]byte{}))}
	copy(ifr.IfrnName[:len(ifr.IfrnName)-1], baseDevice)
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(s), SIOCDELTUNNEL, uintptr(unsafe.Pointer(&ifr))); err != 0 {
		return err
	}
	return nil
}

// Add an Ip address to an interface. This is identical to:
// ip addr add $ip/$ipNet dev $iface
func NetworkLinkAddIpPeer(iface *net.Interface, ip net.IP, peer net.IP, ipNet *net.IPNet) error {
	return networkLinkIpActionPeer(
		syscall.RTM_NEWADDR,
		syscall.NLM_F_CREATE|syscall.NLM_F_EXCL|syscall.NLM_F_ACK,
		IfAddr{iface, ip, ipNet},
		peer,
	)
}

func networkLinkIpActionPeer(action, flags int, ifa IfAddr, peer net.IP) error {
	s, err := getNetlinkSocket()
	if err != nil {
		return err
	}
	defer s.Close()

	family := getIpFamily(ifa.IP)

	wb := newNetlinkRequest(action, flags)

	msg := newIfAddrmsg(family)
	msg.Index = uint32(ifa.Iface.Index)
	prefixLen, _ := ifa.IPNet.Mask.Size()
	msg.Prefixlen = uint8(prefixLen)
	wb.AddData(msg)

	var ipData []byte
	if family == syscall.AF_INET {
		ipData = ifa.IP.To4()
	} else {
		ipData = ifa.IP.To16()
	}

	localData := newRtAttr(syscall.IFA_LOCAL, ipData)
	wb.AddData(localData)

	var peerIpData []byte
	if family == syscall.AF_INET {
		peerIpData = peer.To4()
	} else {
		peerIpData = peer.To16()
	}
	addrData := newRtAttr(syscall.IFA_ADDRESS, peerIpData)
	wb.AddData(addrData)

	if err := s.Send(wb); err != nil {
		return err
	}

	return s.HandleAck(wb.Seq)
}
