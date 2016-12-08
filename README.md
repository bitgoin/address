[![GoDoc](https://godoc.org/github.com/utamaro/base58?status.svg)](https://godoc.org/github.com/utamaro/btcaddr)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/utamaro/btcaddr/LICENSE)


# btcaddr 

## Overview

This  library for handling bitcoin address, including generate private keys from wif, sign/vefiry, and serializing. 

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
	key, err := Generate(BitcoinTest)
	adr := key.PublicKey.Address()
    key2, err := FromWIF(wif, BitcoinTest)
	data := []byte("test data")
	sig, err := private.Sign(data)
	err = key.PublicKey.Verify(sig, data)
...
}
```


# Contribution
Improvements to the codebase and pull requests are encouraged.


