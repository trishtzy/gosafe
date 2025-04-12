package safeutils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// keccak256("SafeMessage(bytes message)");
var SAFE_MESSAGE_TYPEHASH = common.HexToHash("0x60b3cbf8b4a223d68d641b3b6ddf9a298e7f33710cf3d3a9d1146b5a6150fbca")

// HashMessageForSafe hashes a message for a given Safe using its domain serparator
// and returns the signing hash, message hash, or an optional error
func HashMessageForSafe(domainSeparator [32]byte, message interface{}) (common.Hash, common.Hash, error) {
	var messageHash common.Hash
	switch message := message.(type) {
	case string:
		// If the message is a string, we need to hash it using the PersonalSignHash function
		messageHash = PersonalSignHash([]byte(message))
	default:
		return common.Hash{}, common.Hash{}, errors.New("unsupported message type")
	}

	// First encode the message hash as bytes32 - matching abi.encode(bytes32($message_hash))
	bytes32Ty, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		return common.Hash{}, common.Hash{}, fmt.Errorf("failed to create bytes32 type: %v", err)
	}

	// This matches: keccak256(abi.encode(bytes32($message_hash)))
	// We first need to encode messageHash as bytes32
	innerArgs := abi.Arguments{{Type: bytes32Ty}}
	innerPacked, err := innerArgs.Pack(messageHash)
	if err != nil {
		return common.Hash{}, common.Hash{}, fmt.Errorf("failed to pack inner message: %v", err)
	}
	innerHash := crypto.Keccak256Hash(innerPacked)

	// Now pack with SAFE_MESSAGE_TYPEHASH - matching abi.encode(bytes32($safe_type), ...)
	outerArgs := abi.Arguments{
		{Type: bytes32Ty},
		{Type: bytes32Ty},
	}
	outerPacked, err := outerArgs.Pack(SAFE_MESSAGE_TYPEHASH, innerHash)
	if err != nil {
		return common.Hash{}, common.Hash{}, fmt.Errorf("failed to pack outer message: %v", err)
	}

	// This matches: keccak256(abi.encode(bytes32($safe_type), keccak256(abi.encode(bytes32($message_hash)))))
	round1Hash := crypto.Keccak256Hash(outerPacked)

	// Calculate the final hash according to EIP-712
	// This matches: keccak256(abi.encodePacked(bytes1(0x19), bytes1(0x01), bytes32($domain_sep), bytes32($round1)))
	encodedData := []byte{}
	encodedData = append(encodedData, byte(0x19))
	encodedData = append(encodedData, byte(0x01))
	encodedData = append(encodedData, domainSeparator[:]...)
	encodedData = append(encodedData, round1Hash.Bytes()...)

	signingHash := crypto.Keccak256Hash(encodedData)
	return signingHash, messageHash, nil

}

func PersonalSignHash(message []byte) common.Hash {
	signingPrefix := append([]byte("\x19Ethereum Signed Message:\n"), strconv.Itoa(len(message))...)
	message = append(signingPrefix, message...)
	return crypto.Keccak256Hash(message)
}
