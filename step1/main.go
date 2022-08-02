package main

import (
	"fmt"
	"ws/node"
)

func main() {
	yourNode := node.NewBootNode()

	err := yourNode.Start()
	if err != nil {
		return
	}

	fmt.Println(yourNode.GetMultiAddress())
	yourNode.Stop()
}
