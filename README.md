# flashbots ⚡🤖

[![Go Reference](https://pkg.go.dev/badge/github.com/lmittmann/flashbots.svg)](https://pkg.go.dev/github.com/lmittmann/flashbots)

Package flashbots implements RPC API bindings for the Flashbots relay and
[mev-geth](https://github.com/flashbots/mev-geth) for use with the [`w3`](https://github.com/lmittmann/w3)
package.


## Install

```
go get github.com/lmittmann/flashbots
```


## Getting Started

Connect to the Flashbots relay and send a bundle.

```go
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
    // list of signed transactions
	bundle types.Transactions

	bundleHash common.Hash
)
err = client.Call(
	flashbots.SendBundle(&flashbots.SendBundleParam{
		Transactions: bundle,
		BlockNumber:  big.NewInt(999_999_999),
	}).Returns(&bundleHash),
)
if err != nil {
	fmt.Printf("Failed to send bundle: %v\n", err)
	return
}
fmt.Printf("Sent bundle successfully: %s\n", bundleHash)
```

Note that the Flashbots relay does not support batch requests. Thus, sending
more than one request in `Client.Call` will result in a server error.


## RPC Methods

List of supported RPC methods:

Method                     | Go Code
---------------------------|---------
`eth_sendBundle`           | `flashbots.SendBundle(param *flashbots.SendBundleParam).Returns(bundleHash *common.Hash)`
`eth_callBundle`           | TODO <!-- `flashbots.CallBundle(param *flashbots.CallBundleParam).Returns(resp *flashbots.CallBundleResponse)` -->
`flashbots_getUserStats`   | `flashbots.UserStats(blockNumber *big.Int).Returns(resp *flashbots.UserStatsResponse)`
`flashbots_getBundleStats` | `flashbots.BundleStats(bundleHash common.Hash, blockNumber *big.Int).Returns(resp *flashbots.BundleStatsResponse)`
