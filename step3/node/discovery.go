package node

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
)

type DiscoveryPing struct {
	Host host.Host
}

func (d *DiscoveryPing) HandlePeerFound(info peer.AddrInfo) {
	fmt.Println("Found new peer:", info.ID)
}
