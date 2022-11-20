package code

import (
	"fmt"
	"testing"
)

func TestIsValied(t *testing.T) {
	nID, err := NewNatinalCode("3356947567")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(nID.IsValid())
}
