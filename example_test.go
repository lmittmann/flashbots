package flashbots_test

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/flashbots"
	"github.com/lmittmann/w3"
)

func Example() {
	// Private key for request authentication
	var privKey *ecdsa.PrivateKey

	// Connect to relay
	rpcClient, err := rpc.DialHTTPWithClient(
		"https://relay.flashbots.net",
		&http.Client{
			Transport: flashbots.AuthTransport(privKey),
		},
	)
	if err != nil {
		fmt.Printf("Failed to connect to Flashbots relay: %v\n", err)
		return
	}

	client := w3.NewClient(rpcClient)
	defer client.Close()

	// Send bundle
	var (
		bundle types.Transactions // list of signed transactions

		bundleHash common.Hash
	)
	err = client.Call(
		flashbots.SendBundle(&flashbots.SendBundleRequest{
			Transactions: bundle,
			BlockNumber:  big.NewInt(999_999_999),
		}).Returns(&bundleHash),
	)
	if err != nil {
		fmt.Printf("Failed to send bundle to Flashbots relay: %v\n", err)
		return
	}
	fmt.Printf("Sent bundle successfully: %s\n", bundleHash)
}
