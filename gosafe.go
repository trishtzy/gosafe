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

	chainId, err := safe.SafeCaller.GetChainId(nil)
	if err != nil {
		return nil, err
	}

	owners, err := safe.SafeCaller.GetOwners(nil)
	if err != nil {
		return nil, err
	}

	ownersInfo := make([]types.AddressInfo, len(owners))
	for i, owner := range owners {
		ownersInfo[i] = types.AddressInfo{
			Value:   owner.String(),
			Name:    "",
			LogoUri: "",
		}
	}

	nonce, err := safe.SafeCaller.Nonce(nil)
	if err != nil {
		return nil, err
	}

	return &types.Safe{
		Address: types.AddressInfo{
			Value:   config.SafeAddress,
			Name:    "",
			LogoUri: "",
		},
		ChainId: chainId,
		Nonce:   nonce,
		Owners:  ownersInfo,
	}, nil
}
