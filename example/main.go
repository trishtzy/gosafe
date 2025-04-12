package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/trishtzy/gosafe"
)

func main() {
	safeAddress := common.HexToAddress(os.Getenv("SAFE_ADDRESS"))
	client, err := gosafe.New(gosafe.Config{
		SafeAddress:  safeAddress.String(),
		EthClientUrl: "https://mainnet.base.org",
	})
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash("0x1a9469d4f60ff51292be22376fa1717b77d74d4c484f18be4fcafa4f02375beb")
	tx, isPending, err := client.EthClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("isPending: ", isPending)
	fmt.Printf("tx: %+v\n", tx)

	// Get the transaction receipt to get the nonce
	receipt, err := client.EthClient.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	// Get the current nonce of the Safe
	nonce, err := client.Contract.Nonce(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate Safe transaction hash
	safeTxHash, err := client.Contract.GetTransactionHash(
		nil,
		*tx.To(),         // to
		tx.Value(),       // value
		tx.Data(),        // data
		uint8(0),         // operation (0 = CALL)
		big.NewInt(0),    // safeTxGas
		big.NewInt(0),    // baseGas
		big.NewInt(0),    // gasPrice
		common.Address{}, // gasToken
		common.Address{}, // refundReceiver
		nonce,            // nonce
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Safe Transaction Hash: 0x%x\n", safeTxHash)

	// Parse logs from the transaction
	for _, receiptLog := range receipt.Logs {
		// Only try to parse logs from your Safe contract address
		if receiptLog.Address == safeAddress {
			if executionSuccess, err := client.Contract.ParseExecutionSuccess(*receiptLog); err == nil {
				fmt.Printf("ExecutionSuccess event:\n")
				fmt.Printf("TxHash: %x\n", executionSuccess.TxHash)
				fmt.Printf("Payment: %s\n", executionSuccess.Payment)
				continue
			}

			if executionFailure, err := client.Contract.ParseExecutionFailure(*receiptLog); err == nil {
				fmt.Printf("ExecutionFailure event:\n")
				fmt.Printf("TxHash: %x\n", executionFailure.TxHash)
				fmt.Printf("Payment: %s\n", executionFailure.Payment)
				continue
			}

			if moduleSuccess, err := client.Contract.ParseExecutionFromModuleSuccess(*receiptLog); err == nil {
				fmt.Printf("ExecutionFromModuleSuccess event:\n")
				fmt.Printf("Module: %s\n", moduleSuccess.Module.Hex())
				continue
			}

			// Add more event parsers as needed
			fmt.Printf("Error parsing log: %v\n", err)
		}
	}
}
