package native

import "github.com/JeremyOT/sunstone/network"

type nativeManager struct {
}

func New() network.Manager {
	// Work around the fact that LinuxManager doesn't implement Manager
	// when compiled on mac.
	var mgr interface{}
	mgr = &nativeManager{}
	return mgr.(network.Manager)
}
