package flashbots

import (
	"fmt"
	"math/big"
)

type strBigint big.Int

// MarshalJSON implements the json json.Marshaler interface.
func (bi strBigint) MarshalJSON() ([]byte, error) {
	return []byte(`"` + (*big.Int)(&bi).Text(10) + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (bi *strBigint) UnmarshalJSON(input []byte) error {
	if len(input) < 3 || input[0] != '"' || input[len(input)-1] != '"' {
		return fmt.Errorf("invalid number string %s", input)
	}
	_, ok := (*big.Int)(bi).SetString(string(input[1:len(input)-1]), 10)
	if !ok {
		return fmt.Errorf("invalid number string %q", input[1:len(input)-1])
	}
	return nil
}
