package flashbots_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/lmittmann/flashbots"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/rpctest"
)

func TestSendBundle(t *testing.T) {
	rpctest.RunTestCases(t, []rpctest.TestCase[common.Hash]{
		{
			Golden: "send_bundle",
			Call: flashbots.SendBundle(&flashbots.SendBundleRequest{
				RawTransactions: [][]byte{w3.B("0x00"), w3.B("0x01")},
				BlockNumber:     big.NewInt(9_999_999),
				ReplacementUuid: uuid.MustParse("2c9cf5d0-f13c-4b7a-b51d-f462fdb27b51"),
			}),
			WantRet: w3.H("0x2228f5d8954ce31dc1601a8ba264dbd401bf1428388ce88238932815c5d6f23f"),
		},
	})
}
