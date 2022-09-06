package flashbots_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/flashbots"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/rpctest"
)

func TestSendPrivateTx(t *testing.T) {
	tests := []rpctest.TestCase[common.Hash]{
		{
			Golden: "send_private_transaction",
			Call: flashbots.SendPrivateTx(&flashbots.SendPrivateTxRequest{
				RawTx:          w3.B("0x00"),
				MaxBlockNumber: big.NewInt(9_999_999),
				Fast:           true,
			}),
			WantRet: w3.H("0x45df1bc3de765927b053ec029fc9d15d6321945b23cac0614eb0b5e61f3a2f2a"),
		},
	}

	rpctest.RunTestCases(t, tests)
}

func TestCancelPrivateTx(t *testing.T) {
	tests := []rpctest.TestCase[bool]{
		{
			Golden:  "cancel_private_transaction",
			Call:    flashbots.CancelPrivateTx(w3.H("0x45df1bc3de765927b053ec029fc9d15d6321945b23cac0614eb0b5e61f3a2f2a")),
			WantRet: true,
		},
	}

	rpctest.RunTestCases(t, tests)
}
