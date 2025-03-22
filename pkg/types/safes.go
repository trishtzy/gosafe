package types

import "math/big"

type Safe struct {
	Address                    AddressInfo   `json:"address"`
	ChainId                    *big.Int      `json:"chainId"`
	Nonce                      *big.Int      `json:"nonce"`
	Threshold                  int           `json:"threshold"`
	Owners                     []AddressInfo `json:"owners"`
	Implementation             AddressInfo   `json:"implementation"`
	Modules                    []AddressInfo `json:"modules"`
	FallbackHandler            AddressInfo   `json:"fallbackHandler"`
	Guard                      AddressInfo   `json:"guard"`
	Version                    string        `json:"version"`
	ImplementationVersionState string        `json:"implementationVersionState"`
	CollectiblesTag            string        `json:"collectiblesTag"`
	TxQueuedTag                string        `json:"txQueuedTag"`
	TxHistoryTag               string        `json:"txHistoryTag"`
	MessagesTag                string        `json:"messagesTag"`
}

type AddressInfo struct {
	Value   string `json:"value"`
	Name    string `json:"name"`
	LogoUri string `json:"logoUri"`
}
