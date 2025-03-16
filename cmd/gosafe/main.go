package main

import "fmt"

//go:generate abigen --abi ../../internal/abi/contracts/safe.json --pkg contracts --type Safe --out ../../internal/abi/contracts/safe.go
func main() {
	fmt.Println("Hello, World!")
}
