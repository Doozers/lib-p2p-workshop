package node

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
)

type BootNode struct {
	Host host.Host
}

func (b *BootNode) Start() error {
	h, err := libp2p.New()
	if err != nil {
		return err
	}
	b.Host = h
	return nil
}

func (b BootNode) Stop() error {
	err := b.Host.Close()
	if err != nil {
		return err
	}
	return nil
}

func (b BootNode) GetMultiAddress() string {
	return b.Host.Addrs()[2].String() + "/p2p/" + b.Host.ID().String()
}

func NewBootNode() *BootNode {
	return &BootNode{}
}
