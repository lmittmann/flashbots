package flashbots_test

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/flashbots"
)

func Example() {
	// Private key for request signing
	var prv *ecdsa.PrivateKey

	// Connect to Flashbots relay
	client := flashbots.MustDial("https://relay.flashbots.net", prv)
	defer client.Close()

	// Send bundle
	bundle := []*types.Transaction{ /* signed transactions... */ }

	var bundleHash common.Hash
	if err := client.Call(
		flashbots.SendBundle(&flashbots.SendBundleRequest{
			Transactions: bundle,
			BlockNumber:  big.NewInt(999_999_999),
		}).Returns(&bundleHash),
	); err != nil {
		fmt.Printf("Failed to send bundle to Flashbots relay: %v\n", err)
		return
	}
	fmt.Printf("Sent bundle successfully: %s\n", bundleHash)
}
