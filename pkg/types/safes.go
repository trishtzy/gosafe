package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/trishtzy/gosafe/internal/abi/contracts"
)

// OperationType represents the type of operation for a Safe transaction
type OperationType uint8

const (
	Call         OperationType = 0
	DelegateCall OperationType = 1
)

// Safe interacts with Safe contract
type Safe struct {
	Contract    *contracts.Safe
	EthClient   *ethclient.Client
	SafeAddress string
	SafeVersion string
	ChainId     *big.Int
}

// SafeWallet represents a Safe wallet
type SafeWallet struct {
	Address common.Address   `json:"address"`
	ChainId *big.Int         `json:"chainId"`
	Nonce   *big.Int         `json:"nonce"`
	Owners  []common.Address `json:"owners"`
}

// SafeTransaction represents a Safe transaction with its data and signatures
type SafeTransaction struct {
	SafeAddress        common.Address      `json:"safeAddress"`
	ChainID            *big.Int            `json:"chainId"`
	Data               SafeTransactionData `json:"data"`
	Signatures         [][]byte
	ContractSignatures [][]byte // Stores the actual contract signature data
}

// MetaTransactionData represents the base transaction data
type MetaTransactionData struct {
	To        common.Address `json:"to"`
	Value     *big.Int       `json:"value"`
	Data      []byte         `json:"data"`
	Operation OperationType  `json:"operation,omitempty"`
}

// SafeTransactionData represents a complete Safe transaction data structure
type SafeTransactionData struct {
	MetaTransactionData
	SafeTxGas      *big.Int       `json:"safeTxGas"`
	BaseGas        *big.Int       `json:"baseGas"`
	GasPrice       *big.Int       `json:"gasPrice"`
	GasToken       common.Address `json:"gasToken"`
	RefundReceiver common.Address `json:"refundReceiver"`
	Nonce          *big.Int       `json:"nonce"`
}
