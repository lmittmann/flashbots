package flashbots

import (
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/rpctest"
)

func TestCallBundle(t *testing.T) {
	t.Parallel()

	srv := rpctest.NewFileServer(t, "testdata/call_bundle.golden")
	defer srv.Close()

	client := w3.MustDial(srv.URL())
	defer client.Close()

	var (
		callBundleReq = CallBundleRequest{
			RawTransactions:  [][]byte{w3.B("0x00"), w3.B("0x01")},
			BlockNumber:      w3.I("0xb63dcd"),
			StateBlockNumber: nil,
			Timestamp:        big.NewInt(1615920932),
		}
		resp     = new(CallBundleResponse)
		wantResp = &CallBundleResponse{
			BundleGasPrice:    w3.I("476190476193"),
			BundleHash:        w3.H("0x73b1e258c7a42fd0230b2fd05529c5d4b6fcb66c227783f8bece8aeacdd1db2e"),
			CoinbaseDiff:      w3.I("20000000000126000"),
			EthSentToCoinbase: w3.I("20000000000000000"),
			GasFees:           w3.I("126000"), // XXX
			StateBlockNumber:  w3.I("5221585"),
			TotalGasUsed:      42000,
			Results: []CallBundleResult{
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
		}
	)
	if err := client.Call(
		CallBundle(&callBundleReq).Returns(resp),
	); err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	if diff := cmp.Diff(wantResp, resp,
		cmp.AllowUnexported(big.Int{}),
	); diff != "" {
		t.Fatalf("(-want, +got)\n%s", diff)
	}
}
