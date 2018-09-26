package ethconv

import (
	"log"
	"math/big"
	"testing"
)

func TestFromWei(t *testing.T) {
	amount := big.NewInt(1500000000000000000)
	value, err := FromWei(amount, Ether)
	if err != nil {
		t.Error(err)
	}

	if value.Cmp(big.NewFloat(1.5)) != 0 {
		t.Fail()
	}
	log.Println(value)
}

func TestToWei(t *testing.T) {
	amount := big.NewFloat(0.01)
	value, err := ToWei(amount, Ether)
	if err != nil {
		t.Error(err)
	}

	if value.Cmp(big.NewInt(10000000000000000)) != 0 {
		t.Fail()
	}
	log.Println(value)
}
