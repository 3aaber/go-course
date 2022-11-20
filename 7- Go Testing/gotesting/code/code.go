package code

import (
	"errors"
	"fmt"
	"strconv"
)

var posistionValues = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

const lengthOfNationalCode = 10
const modValue = 11

// NationalCode godoc
type NationalCode struct {
	code       [10]int
	stringCode string
	Valid      bool
}

//NewNatinalCode godoc
func NewNatinalCode(code string) (*NationalCode, error) {

	instance := &NationalCode{}

	if len(code) == 10 {

		for key, value := range code {
			digit, err := strconv.Atoi(string(value))
			if err != nil {
				return &NationalCode{}, fmt.Errorf("Problem in parse National ID Number %s in Position %d, value %s", code, key+1, string(value))
			}
			instance.code[key] = digit
		}
		instance.stringCode = code
		return instance, nil
	}
	return &NationalCode{}, errors.New("Wrong Number of Digits")
}

// IsValid godoc
func (n *NationalCode) IsValid() bool {
	var (
		sum   int
		res   int
		index int
	)

	for index = 0; index < lengthOfNationalCode-1; index++ {
		sum = sum + posistionValues[index]*n.code[index]
	}

	res = sum % 11
	controlDigit := n.code[lengthOfNationalCode-1]

	if res < 2 {
		if res == controlDigit {
			n.Valid = true
			return true
		}
	} else {
		if res == 11-controlDigit {
			n.Valid = true
			return true
		}
	}
	n.Valid = false
	return false
}
