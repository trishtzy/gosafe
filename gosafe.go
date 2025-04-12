package gosafe

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/trishtzy/gosafe/internal/abi/contracts"
	"github.com/trishtzy/gosafe/pkg/types"
)

// New creates a new gosafe client
func New(config Config) (*types.Safe, error) {
	ethClient, err := ethclient.Dial(config.EthClientUrl)
	if err != nil {
		return nil, err
	}

	safe, err := contracts.NewSafe(common.HexToAddress(config.SafeAddress), ethClient)
	if err != nil {
		return nil, err
	}

	return &types.Safe{
		Contract:  safe,
		EthClient: ethClient,
	}, nil
}
