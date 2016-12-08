[![GoDoc](https://godoc.org/github.com/utamaro/base58?status.svg)](https://godoc.org/github.com/utamaro/btcaddr)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/utamaro/btcaddr/LICENSE)


# btcaddr 

## Overview

This  library for handling bitcoin address, including generate private keys from wif, sign/vefiry, serializing,
BIP32(Hierarchical Deterministic Bitcoin addresses) and BIP39(mnemonic seed). 

## Requirements

This requires

* git
* go 1.3+


## Installation

     $ go get github.com/utamaro/btcaddr


## Example
(This example omits error handlings for simplicity.)

```go

import btcaddr

func main(){
	key, err := btcaddr.Generate(BitcoinTest)
	adr := key.PublicKey.Address()
    key2, err := btcaddr.FromWIF(wif, BitcoinTest)
	data := []byte("test data")
	sig, err := private.Sign(data)
	err = key.PublicKey.Verify(sig, data)

    seed, err := btcaddr.GenerateSeed(btcaddr.RecommendedSeedLen)
	master, err := btcaddr.NewMasterKey(seed,btcaddr.BitcoinTest)
    derivate,err := master.Child(0)
	priv,err:=derivate.PrivKey()
	derivatepub,err:=derivate.Neuter()
	pub,err:=derivatepub.PubKey()

    entropy, err := btcaddr.NewEntropy(256)
    mnemonic, err := btcaddr.NewMnemonic(entropy)
    seed := btcaddr.NewSeed(mnemonic, "Secret Passphrase")

	..
}
```


# Contribution
Improvements to the codebase and pull requests are encouraged.


