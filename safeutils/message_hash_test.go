package safeutils

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestPersonalSignHash(t *testing.T) {
	message := "hello"
	hash := PersonalSignHash([]byte(message))
	expectedHash := common.HexToHash("0x50b2c43fd39106bafbba0da34fc430e1f91e3c96ea2acee2bc34119f92b37750")
	if hash != expectedHash {
		t.Errorf("Expected hash %v, got %v", expectedHash, hash)
	}
}

func TestHashMessageForSafe(t *testing.T) {
	message := "hello"
	domainSeparator := common.HexToHash("0xcd0f302a4286e706bf130287a133e3491ca6d5c2fb4243cb180378a65da0c04d")
	signingHash, messageHash, err := HashMessageForSafe(domainSeparator, message)
	if err != nil {
		t.Errorf("Error hashing message: %v", err)
	}
	expectedSigningHash := common.HexToHash("0xc1c47b0cbb9bdac2576dc4b84e77dada60020b6ee88520d6e39890e5275e4749")
	expectedMessageHash := common.HexToHash("0x50b2c43fd39106bafbba0da34fc430e1f91e3c96ea2acee2bc34119f92b37750")
	if signingHash != expectedSigningHash {
		t.Errorf("Expected signing hash %v, got %v", expectedSigningHash, signingHash)
	}
	if messageHash != expectedMessageHash {
		t.Errorf("Expected message hash %v, got %v", expectedMessageHash, messageHash)
	}
}
