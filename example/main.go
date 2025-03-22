package main

import (
	"fmt"
	"log"

	"github.com/trishtzy/gosafe"
)

func main() {
	const SafeProxyFactoryAddress = "0x123"
	client, err := gosafe.New(gosafe.Config{
		SafeAddress:  SafeProxyFactoryAddress,
		EthClientUrl: "https://eth.llamarpc.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Address: ", client.Address)
	fmt.Println("ChainId: ", client.ChainId)
	fmt.Println("Nonce: ", client.Nonce)
	fmt.Println("Owners: ", client.Owners)
}
