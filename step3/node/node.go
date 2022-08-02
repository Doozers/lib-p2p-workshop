package node

import "github.com/multiformats/go-multiaddr"

type Node interface {
	Start() error
	Stop() error
	GetMultiAddress() string
	Connect(multiaddr.Multiaddr)
	StartDiscovery()
}
