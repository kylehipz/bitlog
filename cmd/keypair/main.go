package main

import (
	"fmt"
	"github.com/kylehipz/bitlog/pkg/keypair"
)

func main() {
  privateKey, publicKey := keypair.GenerateKeyPair()

  fmt.Println("Here's your private key. Don't share it to anyone")
  fmt.Println()
  fmt.Println(privateKey)

  fmt.Println("Here's your public key. You can share this to anyone")
  fmt.Println()
  fmt.Println(publicKey)
}
