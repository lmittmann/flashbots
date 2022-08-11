package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/lmittmann/flashbots"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

var (
	// random private key
	prv, _ = crypto.GenerateKey()
	// or use the private key you use for signing bundles (and transactions)
	// prv, _ = crypto.HexToECDSA("...")

	// Ethereum mainnet signer
	signer = types.LatestSigner(params.MainnetChainConfig)

	// clients
	client   = w3.MustDial("https://rpc.ankr.com/eth")
	fbClient = flashbots.MustDial("https://relay.flashbots.net", prv)
)

func main() {
	// addr of prv
	addr := crypto.PubkeyToAddress(prv.PublicKey)

	// fetch nonce, gas price, and latest block
	var (
		nonce       uint64
		gasPrice    big.Int
		latestBlock big.Int
	)
	if err := client.Call(
		eth.Nonce(addr, nil).Returns(&nonce),
		eth.GasPrice().Returns(&gasPrice),
		eth.BlockNumber().Returns(&latestBlock),
	); err != nil {
		fmt.Printf("Failed to fetch: %v\n", err)
		return
	}

	// build transaction
	tx := types.MustSignNewTx(prv, signer, &types.DynamicFeeTx{
		Nonce:     nonce,
		GasFeeCap: &gasPrice,
		GasTipCap: w3.I("1 gwei"),
		Gas:       250_000,
		// To:     w3.APtr("0x..."),
		// Data:   w3.B("0xc0fe..."),
	})

	// call bundle
	var callBundle flashbots.CallBundleResponse
	if err := fbClient.Call(
		flashbots.CallBundle(&flashbots.CallBundleRequest{
			Transactions: []*types.Transaction{tx},
			BlockNumber:  new(big.Int).Add(&latestBlock, w3.Big1),
		}).Returns(&callBundle),
	); err != nil {
		fmt.Printf("Failed to call bundle: %v\n", err)
		return
	}
	fmt.Printf("Call bundle response: %+v\n", callBundle)

	// send bundle
	var bundleHash common.Hash
	if err := fbClient.Call(
		flashbots.SendBundle(&flashbots.SendBundleRequest{
			Transactions: []*types.Transaction{tx},
			BlockNumber:  new(big.Int).Add(&latestBlock, w3.Big1),
		}).Returns(&bundleHash),
	); err != nil {
		fmt.Printf("Failed to send bundle: %v\n", err)
		return
	}
	fmt.Printf("Bundle hash: %s\n", bundleHash)
}
