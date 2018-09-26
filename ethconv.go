package ethconv

import (
	"fmt"
	"math/big"
)

// FromWei convert from wei to a unit
func FromWei(amount *big.Int, unit string) (*big.Float, error) {
	result := new(big.Float)
	result.SetInt(amount)
	if unit == Wei {
		return result, nil
	}
	value, err := getUnitValue(unit)
	if err != nil {
		return nil, err
	}
	return result.Quo(result, new(big.Float).SetInt(value)), nil
}

// ToWei convert a unit to wei
func ToWei(amount *big.Float, unit string) (*big.Int, error) {
	if unit == Wei {
		return bigFloatToBigInt(amount), nil
	}
	value, err := getUnitValue(unit)
	if err != nil {
		return nil, err
	}
	amount.Mul(amount, new(big.Float).SetInt(value))
	result := new(big.Int)
	amount.Int(result)
	return result, nil
}

// getUnitValue from Units map
func getUnitValue(unit string) (*big.Int, error) {
	value, ok := Units[unit]
	if !ok {
		return nil, fmt.Errorf("unit %s is not supported", unit)
	}
	result := new(big.Int)
	result.SetString(value, 10)
	return result, nil
}

// bigFloatToBigInt conversion
func bigFloatToBigInt(value *big.Float) *big.Int {
	per := new(big.Float)
	per.SetInt(big.NewInt(1000000000000000000))
	value.Mul(value, per)
	result := new(big.Int)
	value.Int(result)
	return result
}
