package flashbots

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/rpctest"
)

func TestSendPrivateTransaction(t *testing.T) {
	t.Parallel()

	srv := rpctest.NewFileServer(t, "testdata/send_private_transaction.golden")
	defer srv.Close()

	client := w3.MustDial(srv.URL())
	defer client.Close()

	var (
		sendPrivateTransactionReq = &SendPrivateTransactionRequest{
			RawTransaction: w3.B("0x00"),
			MaxBlockNumber: big.NewInt(9_999_999),
			Fast:           true,
		}
		hash     common.Hash
		wantHash = w3.H("0x45df1bc3de765927b053ec029fc9d15d6321945b23cac0614eb0b5e61f3a2f2a")
	)

	if err := client.Call(
		SendPrivateTransaction(sendPrivateTransactionReq).Returns(&hash),
	); err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	if wantHash != hash {
		t.Fatalf("want %v, got %v", wantHash, hash)
	}
}
