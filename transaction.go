package gosafe

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/trishtzy/gosafe/pkg/types"
)

// GetTransactionHash returns the hash of a Safe transaction
func (s *types.Safe) GetTransactionHash(safeTransaction *types.SafeTransaction) (common.Hash, error) {
	safeAddress := s.SafeAddress
	safeVersion := s.SafeVersion
	chainId := s.ChainId

	return calculateSafeTransactionHash(safeAddress, &safeTransaction.Data, safeVersion, chainId)
}

// calculateSafeTransactionHash calculates the hash of a Safe transaction
func calculateSafeTransactionHash(safeAddress string, safeTransaction *types.SafeTransactionData, safeVersion *big.Int, chainId *big.Int) (common.Hash, error) {
	return common.Hash{}, nil
}
