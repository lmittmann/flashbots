package internal

import (
	"bytes"
	"errors"
	"math/big"
	"strconv"
	"testing"
)

func TestStrBigintMarshalJSON(t *testing.T) {
	tests := []struct {
		Int      StrInt
		WantJSON []byte
		WantErr  error
	}{
		{
			Int:      StrInt(*big.NewInt(0)),
			WantJSON: []byte(`"0"`),
		},
		{
			Int:      StrInt(*big.NewInt(1)),
			WantJSON: []byte(`"1"`),
		},
		{
			Int:      StrInt(*big.NewInt(-1)),
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
		WantInt StrInt
		WantErr error
	}{
		{
			JSON:    []byte(`""`),
			WantInt: StrInt(*big.NewInt(0)),
		},
		{
			JSON:    []byte(`"0"`),
			WantInt: StrInt(*big.NewInt(0)),
		},
		{
			JSON:    []byte(`"1"`),
			WantInt: StrInt(*big.NewInt(1)),
		},
		{
			JSON:    []byte(`"-1"`),
			WantInt: StrInt(*big.NewInt(-1)),
		},
		{
			JSON:    []byte(`0`),
			WantErr: errors.New("invalid number string 0"),
		},
		{
			JSON:    []byte(`"xxx"`),
			WantErr: errors.New(`invalid number string "xxx"`),
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotInt := new(StrInt)
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
