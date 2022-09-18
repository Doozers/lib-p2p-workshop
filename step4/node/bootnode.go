package node

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

type BootNode struct {
	Host host.Host
	mdns *mdns.Service
}

func (b *BootNode) Start() error {
	h, err := libp2p.New()
	if err != nil {
		return err
	}
	b.Host = h
	return nil
}

func (b *BootNode) Stop() error {
	if err := b.Host.Close(); err != nil {
		return err
	}
	return nil
}

func (b *BootNode) GetMultiAddress() (string, error) {
	if len(b.Host.Addrs()) < 3 {
		return "", fmt.Errorf("no multiaddress available")
	}
	return b.Host.Addrs()[2].String() + "/p2p/" + b.Host.ID().String(), nil
}

func (b *BootNode) Connect(stringAddr string) {
	addr, err := peer.AddrInfoFromString(stringAddr)
	if err != nil {
		return
	}

	if err = b.Host.Connect(context.Background(), *addr); err != nil {
		return
	}

	fmt.Println("Connected to", addr.ID)
	return
}

func (b *BootNode) StartDiscovery() error {
	service := mdns.NewMdnsService(b.Host, "peer-discovery", &DiscoveryPing{Host: b.Host})
	err := service.Start()
	if err != nil {
		return err
	}

	return nil
}

func NewBootNode() *BootNode {
	return &BootNode{}
}
