package flashbots_test

import (
	"math/big"
	"testing"

	"github.com/lmittmann/flashbots"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/rpctest"
)

func TestCallBundle(t *testing.T) {
	tests := []rpctest.TestCase[flashbots.CallBundleResponse]{
		{
			Golden: "call_bundle",
			Call: flashbots.CallBundle(&flashbots.CallBundleRequest{
				RawTransactions:  [][]byte{w3.B("0x00"), w3.B("0x01")},
				BlockNumber:      w3.I("0xb63dcd"),
				StateBlockNumber: nil,
				Timestamp:        big.NewInt(1615920932),
			}),
			WantRet: flashbots.CallBundleResponse{
				BundleGasPrice:    w3.I("476190476193"),
				BundleHash:        w3.H("0x73b1e258c7a42fd0230b2fd05529c5d4b6fcb66c227783f8bece8aeacdd1db2e"),
				CoinbaseDiff:      w3.I("20000000000126000"),
				EthSentToCoinbase: w3.I("20000000000000000"),
				GasFees:           w3.I("126000"), // XXX
				StateBlockNumber:  w3.I("5221585"),
				TotalGasUsed:      42000,
				Results: []flashbots.CallBundleResult{
					{
						CoinbaseDiff:      w3.I("10000000000063000"),
						EthSentToCoinbase: w3.I("10000000000000000"),
						FromAddress:       w3.A("0x02A727155aeF8609c9f7F2179b2a1f560B39F5A0"),
						GasFees:           w3.I("63000"), //XXX
						GasPrice:          w3.I("476190476193"),
						GasUsed:           21000,
						ToAddress:         w3.APtr("0x73625f59CAdc5009Cb458B751b3E7b6b48C06f2C"),
						TxHash:            w3.H("0x669b4704a7d993a946cdd6e2f95233f308ce0c4649d2e04944e8299efcaa098a"),
						Value:             w3.B("0x"),
					},
					{
						CoinbaseDiff:      w3.I("10000000000063000"),
						EthSentToCoinbase: w3.I("10000000000000000"),
						FromAddress:       w3.A("0x02A727155aeF8609c9f7F2179b2a1f560B39F5A0"),
						GasFees:           w3.I("63000"),
						GasPrice:          w3.I("476190476193"),
						GasUsed:           21000,
						ToAddress:         w3.APtr("0x73625f59CAdc5009Cb458B751b3E7b6b48C06f2C"),
						TxHash:            w3.H("0xa839ee83465657cac01adc1d50d96c1b586ed498120a84a64749c0034b4f19fa"),
						Value:             w3.B("0x"),
					},
				},
			},
		},
	}

	rpctest.RunTestCases(t, tests)
}
