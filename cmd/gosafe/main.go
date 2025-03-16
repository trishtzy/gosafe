package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/trishtzy/gosafe/internal/abi/contracts"
)

const SAFE_PROXY_FACTORY_ADDRESS = "0x66b3428018ca07e41746f8a8fad8bf7b37238f0a"

//go:generate abigen --abi ../../internal/abi/contracts/safe.json --pkg contracts --type Safe --out ../../internal/abi/contracts/safe.go
func main() {
	// Connect to an Ethereum node
	client, err := ethclient.Dial("https://eth.llamarpc.com") // e.g. Infura URL, local node, etc.
	if err != nil {
		panic(err)
	}

	safe, err := contracts.NewSafe(common.HexToAddress(SAFE_PROXY_FACTORY_ADDRESS), client)
	if err != nil {
		panic(err)
	}

	chainId, err := safe.SafeCaller.GetChainId(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(chainId)
}
