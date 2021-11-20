package keypair

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

func encodeKeys(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
  x509EncodedPrivate, _ := x509.MarshalECPrivateKey(privateKey)

  privatePemBlock := pem.Block{
    Type: "PRIVATE KEY",
    Bytes: x509EncodedPrivate,
  }

  privateEncodedPem := pem.EncodeToMemory(&privatePemBlock)

  x509EncodedPublic, _ := x509.MarshalPKIXPublicKey(publicKey)

  publicPemBlock := pem.Block{
    Type: "PUBLIC KEY",
    Bytes: x509EncodedPublic,
  }
  
  publicEncodedPem := pem.EncodeToMemory(&publicPemBlock)

  return string(privateEncodedPem), string(publicEncodedPem)
}
