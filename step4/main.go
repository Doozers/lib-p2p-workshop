package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"ws/node"
)

func main() {

	yourNode := node.NewBootNode()

	err := yourNode.Start()
	if err != nil {
		return
	}
	defer yourNode.Stop()

	fmt.Println(yourNode.GetMultiAddress())

	if len(os.Args) > 1 {
		yourNode.Connect(os.Args[1])
	}

	err = yourNode.StartDiscovery()
	if err != nil {
		fmt.Println(err)
		return
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
