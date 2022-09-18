package node

import (
	"fmt"

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

func NewBootNode() *BootNode {
	return &BootNode{}
}
