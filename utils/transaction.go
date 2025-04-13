package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/trishtzy/gosafe/pkg/types"
)

// GetTransactionHash returns the hash of a Safe transaction
// Currently only supports Safe v1.3.0 and above
func GetTransactionHash(safeTransaction *types.SafeTransaction) []byte {
	// Pack the transaction data
	safeTxHash := crypto.Keccak256(encode(safeTransaction))

	// Create domain separator
	domainSeparator := getDomainSeparator(safeTransaction)

	// Final hash
	var message []byte
	// 0x19: Marks the start of an Ethereum signed message
	// 0x01: Version byte indicating EIP-712 structured data
	message = append(message, []byte{0x19, 0x01}...)
	message = append(message, domainSeparator...)
	message = append(message, safeTxHash...)

	return crypto.Keccak256(message)
}

// encode packs the Safe transaction data for hashing
func encode(safeTransaction *types.SafeTransaction) []byte {
	// Pack according to EIP-712 structure
	var data []byte

	// Encode SafeTx type hash
	typeHash := crypto.Keccak256([]byte("SafeTx(address to,uint256 value,bytes data,uint8 operation,uint256 safeTxGas,uint256 baseGas,uint256 gasPrice,address gasToken,address refundReceiver,uint256 nonce)"))
	data = append(data, typeHash...)

	// Pack transaction parameters
	data = append(data, common.LeftPadBytes(safeTransaction.Data.To.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(safeTransaction.Data.Value.Bytes(), 32)...)
	data = append(data, crypto.Keccak256(safeTransaction.Data.Data)...)
	data = append(data, common.LeftPadBytes([]byte{byte(safeTransaction.Data.Operation)}, 32)...)
	data = append(data, common.LeftPadBytes(safeTransaction.Data.SafeTxGas.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(safeTransaction.Data.BaseGas.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(safeTransaction.Data.GasPrice.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(safeTransaction.Data.GasToken.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(safeTransaction.Data.RefundReceiver.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(safeTransaction.Data.Nonce.Bytes(), 32)...)

	return data
}

// getDomainSeparator returns the EIP-712 domain separator
func getDomainSeparator(safeTransaction *types.SafeTransaction) []byte {
	domainTypeHash := crypto.Keccak256([]byte("EIP712Domain(uint256 chainId,address verifyingContract)"))

	var domain []byte
	domain = append(domain, domainTypeHash...)
	domain = append(domain, common.LeftPadBytes(safeTransaction.ChainID.Bytes(), 32)...)
	domain = append(domain, common.LeftPadBytes(safeTransaction.SafeAddress.Bytes(), 32)...)

	return crypto.Keccak256(domain)
}
