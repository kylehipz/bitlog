package keypair

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

func GenerateKeyPair() (string, string) {
  privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
  publicKey := &privateKey.PublicKey

  if err != nil {
    log.Fatal(err.Error())
  }

  return encodeKeys(privateKey, publicKey)
}
