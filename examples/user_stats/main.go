package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/lmittmann/flashbots"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

var (
	// random private key
	prv, _ = crypto.GenerateKey()
	// or use the private key you use for signing bundles
	// prv, _ = crypto.HexToECDSA("...")

	client   = w3.MustDial("https://rpc.ankr.com/eth")
	fbClient = flashbots.MustDial("https://relay.flashbots.net", prv)
)

func main() {
	// fetch latest block
	var latestBlock big.Int
	if err := client.Call(
		eth.BlockNumber().Returns(&latestBlock),
	); err != nil {
		fmt.Printf("Failed to fetch latest block: %v\n", err)
		return
	}

	// fetch user statistics
	var userStats flashbots.UserStatsResponse
	if err := fbClient.Call(
		flashbots.UserStats(&latestBlock).Returns(&userStats),
	); err != nil {
		fmt.Printf("Faile to fetch user statistics: %v\n", err)
	}

	// print user statistics
	fmt.Printf("High priority: %t\n", userStats.IsHighPriority)
	fmt.Printf("7 day fees: %s ETH\n", w3.FromWei(userStats.Last7dMinerPayments, 18))
	fmt.Printf("Total fees: %s ETH\n", w3.FromWei(userStats.AllTimeMinerPayments, 18))
}
