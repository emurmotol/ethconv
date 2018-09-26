# Ethereum Unit Converter 

A simple package for performing easy conversion between various Ethereum units.

The conversion logic has been modelled after [web3.js](https://github.com/ethereum/web3.js)

This package is a golang version of [ethereum-converter](https://github.com/mbezhanov/ethereum-converter) PHP library.

**Installation:**

```bash
go get github.com/emurmotol/ethconv
```

**Usage:**

```go
package main

import (
	"fmt"
    "math/big"
    "github.com/emurmotol/ethconv"
)

func main() {
    // wei to ether
    fromWei, err := ethconv.FromWei(big.NewInt(1500000000000000000), ethconv.Ether)
    if err != nil {
        panic(err)
    }
    fmt.Println(fromWei) // 1.5

    // ether to wei
    toWei, err := ethconv.ToWei(big.NewFloat(0.01), ethconv.Ether)
    if err != nil {
        panic(err)
    }
    fmt.Println(toWei) // 10000000000000000
}
```

**Supported units:**

- `kwei` / `ada`
- `mwei` / `babbage`
- `gwei` / `shannon`
- `szabo`
- `finney`
- `ether`
- `kether` / `grand` / `einstein`
- `mether`
- `gether`
- `tether`