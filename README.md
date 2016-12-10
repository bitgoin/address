[![GoDoc](https://godoc.org/github.com/bitgoin/address?status.svg)](https://godoc.org/github.com/bitgoin/address)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/bitgoin/address/LICENSE)


# btcaddr 

## Overview

This  library is for handling bitcoin address, including generate private keys from wif, sign/vefiry, serializing,
BIP32(Hierarchical Deterministic Bitcoin addresses) and BIP39(mnemonic seed). 

## Requirements

This requires

* git
* go 1.3+


## Installation

     $ go get github.com/bitgoin/address


## Example
(This example omits error handlings for simplicity.)

```go

import address

func main(){
	key, err := address.Generate(BitcoinTest)
	adr := key.PublicKey.Address()
    key2, err := address.FromWIF(wif, BitcoinTest)
	data := []byte("test data")
	sig, err := private.Sign(data)
	err = key.PublicKey.Verify(sig, data)

    seed, err := address.GenerateSeed(address.RecommendedSeedLen)
	master, err := address.NewMasterKey(seed,address.BitcoinTest)
    derivate,err := master.Child(0)
	priv,err:=derivate.PrivKey()
	derivatepub,err:=derivate.Neuter()
	pub,err:=derivatepub.PubKey()

    entropy, err := address.NewEntropy(256)
    mnemonic, err := address.NewMnemonic(entropy)
    seed := address.NewSeed(mnemonic, "Secret Passphrase")

	..
}
```


# Contribution
Improvements to the codebase and pull requests are encouraged.


