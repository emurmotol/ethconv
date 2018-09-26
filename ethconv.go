package ethconv

import (
	"math/big"
)

func FromWei(amount *big.Int, unit string) (*big.Int, bool) {
	if unit == "wei" {
		return amount, true
	}
	value, ok := getUnitValue(unit)
	if !ok {
		return nil, false
	}
	return value.Div(amount, amount), true
}

func ToWei(amount *big.Int, unit string) (*big.Int, bool) {
	if unit == "wei" {
		return amount, true
	}
	value, ok := getUnitValue(unit)
	if !ok {
		return nil, false
	}
	return value.Div(amount, amount), true
}

func getUnitValue(unit string) (*big.Int, bool) {
	target := new(big.Int)
	value, ok := Units[unit]
	if !ok {
		return nil, false
	}
	return target.SetString(value, 10)
}
