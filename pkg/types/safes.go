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

// SafeTransaction represents a Safe transaction with its data and signatures
type SafeTransaction struct {
	Data       SafeTransactionData       `json:"data"`
	Signatures map[string]*SafeSignature `json:"signatures"`
}

// MetaTransactionData represents the base transaction data
type MetaTransactionData struct {
	To        string        `json:"to"`
	Value     string        `json:"value"`
	Data      string        `json:"data"`
	Operation OperationType `json:"operation,omitempty"`
}

// SafeTransactionData represents a complete Safe transaction data structure
type SafeTransactionData struct {
	MetaTransactionData
	SafeTxGas      string `json:"safeTxGas"`
	BaseGas        string `json:"baseGas"`
	GasPrice       string `json:"gasPrice"`
	GasToken       string `json:"gasToken"`
	RefundReceiver string `json:"refundReceiver"`
	Nonce          uint64 `json:"nonce"`
}

// SafeSignature represents a signature for a Safe transaction
type SafeSignature struct {
	Signer              string `json:"signer"`
	Data                string `json:"data"`
	IsContractSignature bool   `json:"isContractSignature"`
}
