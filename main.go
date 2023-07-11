package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/secp256k1/ecdsa"
)

func main() {
	// generate parameters
	privKey, _ := ecdsa.GenerateKey(rand.Reader)
	publicKey := privKey.PublicKey

	// sign
	msg := []byte("testing ECDSA (sha256)")
	md := sha256.New()
	_, _ = md.Write(msg)
	hashedMsg := md.Sum(nil)
	sigBin, _ := privKey.Sign(hashedMsg, nil)

	// check that the signature is correct
	flag, _ := publicKey.Verify(sigBin, hashedMsg, nil)
	if !flag {
		fmt.Println("can't verify signature")
	} else {
		fmt.Println("signature verification ok")
	}

	// // unmarshal signature
	var sig ecdsa.Signature
	sig.SetBytes(sigBin)
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(sig.R[:32])
	s.SetBytes(sig.S[:32])

	fmt.Printf("hashedMsg: %x\n", hashedMsg)
	fmt.Printf("After HashToInt(): %x\n", ecdsa.HashToInt(hashedMsg))
	fmt.Printf("Sig.R: %x\n", r)
	fmt.Printf("Sig.S: %x\n", s)
	fmt.Printf("Pub.X: %x\n", privKey.PublicKey.A.X.Bytes())
	fmt.Printf("Pub.Y: %x\n", privKey.PublicKey.A.Y.Bytes())
}

