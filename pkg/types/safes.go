package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/trishtzy/gosafe/internal/abi/contracts"
)

// Safe interacts with Safe contract
type Safe struct {
	Contract  *contracts.Safe
	EthClient *ethclient.Client
}

// SafeWallet represents a Safe wallet
type SafeWallet struct {
	Address common.Address   `json:"address"`
	ChainId *big.Int         `json:"chainId"`
	Nonce   *big.Int         `json:"nonce"`
	Owners  []common.Address `json:"owners"`
}
