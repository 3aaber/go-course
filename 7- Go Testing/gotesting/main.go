package main

import (
	"fmt"
	"ncode/code"
)

func main() {

	nID, err := code.NewNatinalCode("3356947567")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(nID.IsValid())
	fmt.Println(nID.Valid)
}
