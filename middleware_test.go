package flashbots

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestSign(t *testing.T) {
	t.Parallel()

	privKey, err := crypto.HexToECDSA("0000000000000000000000000000000000000000000000000000000000000001")
	if err != nil {
		t.Fatalf("Failed to read key: %v", err)
	}

	authRT := &authRoundTripper{
		privKey: privKey,
		addr:    crypto.PubkeyToAddress(privKey.PublicKey),
	}
	body := []byte(`{"jsonrpc":"2.0","id":1,"method":"eth_sendBundle","params":[{"txs":["0x00","0x01"],"blockNumber":"0x98967f"}]}`)
	gotSig, err := authRT.sign(body)
	if err != nil {
		t.Fatalf("Failed to sign body: %v", err)
	}

	wantSig := "0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf:0x2765bcbc32f0c6fc822e1d34e188f8337ec52524a7fd4346ba3ca785f3c641a51aaabe9b9392657ab0fd635fb0b527b2dacca7fea1b6b1c3eae553ded693073e01"
	if wantSig != gotSig {
		t.Fatalf("want %s\ngot  %s", wantSig, gotSig)
	}
}

func BenchmarkSign(b *testing.B) {
	privKey, err := crypto.HexToECDSA("0000000000000000000000000000000000000000000000000000000000000001")
	if err != nil {
		b.Fatalf("Failed to read key: %v", err)
	}

	authRT := &authRoundTripper{
		privKey: privKey,
		addr:    crypto.PubkeyToAddress(privKey.PublicKey),
	}
	body := []byte(`{"jsonrpc":"2.0","id":1,"method":"eth_sendBundle","params":[{"txs":["0x00","0x01"],"blockNumber":"0x98967f"}]}`)

	for i := 0; i < b.N; i++ {
		if _, err := authRT.sign(body); err != nil {
			b.Fatalf("Faild to sign body: %v", err)
		}
	}
}
