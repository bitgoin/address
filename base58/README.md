[![GoDoc](https://godoc.org/github.com/utamaro/base58?status.svg)](https://godoc.org/github.com/utamaro/base58)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/utamaro/base58/LICENSE)


# base58 

## Overview

This  is [base58check](https://en.bitcoin.it/wiki/Base58Check_encoding) library,
from https://github.com/prettymuchbryce/hellobitcoin and https://github.com/tv42/base58.

That's it.

## Requirements

This requires

* git
* go 1.3+


## Installation

     $ go get github.com/utamaro/base58


## Example
(This example omits error handlings for simplicity.)

```go

import base58

func main(){
	test:=[]byte{0x01,0x02}
	result, err := Sum(test)

...
}
```


# Contribution
Improvements to the codebase and pull requests are encouraged.


