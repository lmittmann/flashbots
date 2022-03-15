package internal

import (
	"fmt"
	"math/big"
)

// StrInt wraps a big.Int and is marshaled as a string.
type StrInt big.Int

// MarshalJSON implements the json json.Marshaler interface.
func (i StrInt) MarshalJSON() ([]byte, error) {
	return []byte(`"` + (*big.Int)(&i).Text(10) + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *StrInt) UnmarshalJSON(input []byte) error {
	if len(input) < 3 || input[0] != '"' || input[len(input)-1] != '"' {
		return fmt.Errorf("invalid number string %s", input)
	}
	_, ok := (*big.Int)(i).SetString(string(input[1:len(input)-1]), 10)
	if !ok {
		return fmt.Errorf("invalid number string %q", input[1:len(input)-1])
	}
	return nil
}
