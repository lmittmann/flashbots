package flashbots

import (
	"bytes"
	"math/big"
	"strconv"
	"testing"
)

func TestStrBigintMarshalJSON(t *testing.T) {
	tests := []struct {
		Int      strBigint
		WantJSON []byte
		WantErr  error
	}{
		{
			Int:      strBigint(*big.NewInt(0)),
			WantJSON: []byte(`"0"`),
		},
		{
			Int:      strBigint(*big.NewInt(1)),
			WantJSON: []byte(`"1"`),
		},
		{
			Int:      strBigint(*big.NewInt(-1)),
			WantJSON: []byte(`"-1"`),
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotJSON, err := test.Int.MarshalJSON()
			if err != nil {
				if test.WantErr == nil {
					t.Fatalf("Unexpected error: %v", err)
				} else if err.Error() != test.WantErr.Error() {
					t.Fatalf("Unexpected error:\nwant %s\ngot  %s", test.WantErr, err)
				}
			}
			if !bytes.Equal(test.WantJSON, gotJSON) {
				t.Fatalf("Wrong JSON:\nwant %s\ngot  %s", test.WantJSON, gotJSON)
			}
		})
	}
}
func TestStrBigintUnmarshalJSON(t *testing.T) {
	tests := []struct {
		JSON    []byte
		WantInt strBigint
		WantErr error
	}{
		{
			JSON:    []byte(`"0"`),
			WantInt: strBigint(*big.NewInt(0)),
		},
		{
			JSON:    []byte(`"1"`),
			WantInt: strBigint(*big.NewInt(1)),
		},
		{
			JSON:    []byte(`"-1"`),
			WantInt: strBigint(*big.NewInt(-1)),
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotInt := new(strBigint)
			err := gotInt.UnmarshalJSON(test.JSON)
			if err != nil {
				if test.WantErr == nil {
					t.Fatalf("Unexpected error: %v", err)
				} else if err.Error() != test.WantErr.Error() {
					t.Fatalf("Unexpected error:\nwant %s\ngot  %s", test.WantErr, err)
				}
			}
			if (*big.Int)(gotInt).Cmp((*big.Int)(&test.WantInt)) != 0 {
				t.Fatalf("Wrong Int:\nwant %v\ngot  %v", test.WantInt, gotInt)
			}
		})
	}
}
