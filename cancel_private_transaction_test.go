package flashbots

import (
	"testing"

	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/rpctest"
)

func TestCancelPrivateTransaction(t *testing.T) {
	t.Parallel()

	srv := rpctest.NewFileServer(t, "testdata/cancel_private_transaction.golden")
	defer srv.Close()

	client := w3.MustDial(srv.URL())
	defer client.Close()

	var (
		success     bool
		wantSuccess = true
	)
	if err := client.Call(
		CancelPrivateTx(w3.H("0x45df1bc3de765927b053ec029fc9d15d6321945b23cac0614eb0b5e61f3a2f2a")).Returns(&success),
	); err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	if wantSuccess != success {
		t.Fatalf("want %v, got %v", wantSuccess, success)
	}
}
