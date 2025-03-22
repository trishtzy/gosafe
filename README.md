# gosafe

go sdk for safe-core-sdk

## Example

```go
const SafeProxyFactoryAddress = "0x66b3428018ca07e41746f8a8fad8bf7b37238f0a"
client, err := gosafe.New(gosafe.Config{
  SafeAddress:  SafeProxyFactoryAddress,
  EthClientUrl: "https://eth.llamarpc.com",
})
if err != nil {
  log.Fatal(err)
}

fmt.Println("Address: ", client.Address)
fmt.Println("ChainId: ", client.ChainId)
fmt.Println("Nonce: ", client.Nonce)
fmt.Println("Owners: ", client.Owners)
```
