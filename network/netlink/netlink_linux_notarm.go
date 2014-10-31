// +build !arm

// From https://github.com/docker/libcontainer

package netlink

func ifrDataByte(b byte) int8 {
	return int8(b)
}
