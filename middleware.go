package flashbots

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// AuthTransport returns a http.RoundTripper that adds the
// 'X-Flashbots-Signature' header to every request.
func AuthTransport(privKey *ecdsa.PrivateKey) http.RoundTripper {
	if privKey == nil {
		return &authRoundTripper{}
	}
	return &authRoundTripper{privKey, crypto.PubkeyToAddress(privKey.PublicKey), http.DefaultTransport}
}

type authRoundTripper struct {
	privKey *ecdsa.PrivateKey
	addr    common.Address
	next    http.RoundTripper
}

func (auth *authRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if auth.privKey == nil {
		return nil, errors.New("flashbots: key is nil")
	}

	if r.Body != nil {
		// write request body to buffer and set buffer as new body
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, r.Body); err != nil {
			return nil, err
		}
		r.Body.Close()
		r.Body = ioutil.NopCloser(buf)

		// generate payload signature
		sig, err := auth.sign(buf.Bytes())
		if err != nil {
			return nil, err
		}
		r.Header.Set("X-Flashbots-Signature", sig)
	}
	return auth.next.RoundTrip(r)
}

func (auth *authRoundTripper) sign(body []byte) (string, error) {
	bodyHash := crypto.Keccak256(body)
	sig, err := crypto.Sign(accounts.TextHash([]byte(hexutil.Encode(bodyHash))), auth.privKey)
	if err != nil {
		return "", err
	}
	return auth.addr.Hex() + ":" + hexutil.Encode(sig), nil
}
