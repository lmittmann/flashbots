package flashbots

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/rpctest"
)

func TestSendBundle(t *testing.T) {
	t.Parallel()

	srv := rpctest.NewFileServer(t, "testdata/send_bundle.golden")
	defer srv.Close()

	client := w3.MustDial(srv.URL())
	defer client.Close()

	var (
		sendBundleParam = SendBundleParam{
			RawTransactions: [][]byte{w3.B("0x00"), w3.B("0x01")},
			BlockNumber:     big.NewInt(9_999_999),
		}
		hash     common.Hash
		wantHash = w3.H("0x2228f5d8954ce31dc1601a8ba264dbd401bf1428388ce88238932815c5d6f23f")
	)

	if err := client.Call(SendBundle(&sendBundleParam).Returns(&hash)); err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	if wantHash != hash {
		t.Fatalf("want %v, got %v", wantHash, hash)
	}
}
