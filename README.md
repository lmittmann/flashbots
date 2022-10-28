# flashbots ⚡🤖

[![Go Reference](https://pkg.go.dev/badge/github.com/lmittmann/flashbots.svg)](https://pkg.go.dev/github.com/lmittmann/flashbots)
[![Go Report Card](https://goreportcard.com/badge/github.com/lmittmann/flashbots)](https://goreportcard.com/report/github.com/lmittmann/flashbots)
[![Coverage Status](https://coveralls.io/repos/github/lmittmann/flashbots/badge.svg?branch=main)](https://coveralls.io/github/lmittmann/flashbots?branch=main)
[![Latest Release](https://img.shields.io/github/v/release/lmittmann/flashbots)](https://github.com/lmittmann/flashbots/releases)

Package flashbots implements RPC API bindings for the Flashbots relay and
[mev-geth](https://github.com/flashbots/mev-geth) for use with the [`w3` package](https://github.com/lmittmann/w3).


## Install

```
go get github.com/lmittmann/flashbots
```


## Getting Started

> **Note**
> Check out the [examples](examples/)!

Connect to the Flashbots relay. The [`w3.Client`](https://pkg.go.dev/github.com/lmittmann/w3#Client)
returned by [`Dial`](https://pkg.go.dev/github.com/lmittmann/flashbots#Dial)
uses the [`AuthTransport`](https://pkg.go.dev/github.com/lmittmann/flashbots#AuthTransport)
to add the `X-Flashbots-Signature` header to every request.

```go
// Private key for request signing.
var prv *ecdsa.PrivateKey

// Connect to Flashbots Relay
client := flashbots.MustDial("https://relay.flashbots.net", prv)
defer client.Close()
```

Send a bundle to the Flashbots relay.

```go
bundle := []*types.Transaction{ /* signed transactions... */ }

var bundleHash common.Hash
err := client.Call(
	flashbots.SendBundle(&flashbots.SendBundleRequest{
		Transactions: bundle,
		BlockNumber:  big.NewInt(999_999_999),
	}).Returns(&bundleHash),
)
```

> **Warning**
> The Flashbots relay does not support batch requests. Thus, sending more than
one call in `Client.Call` will result in a server error.


## RPC Methods

List of supported RPC methods.

| Method                         | Go Code
| :----------------------------- | :-------
| `eth_sendBundle`               | `flashbots.SendBundle(r *flashbots.SendBundleRequest).Returns(bundleHash *common.Hash)`
| `eth_callBundle`               | `flashbots.CallBundle(r *flashbots.CallBundleRequest).Returns(resp *flashbots.CallBundleResponse)`
| `eth_sendPrivateTransaction`   | `flashbots.SendPrivateTx(r *flashbots.SendPrivateTxRequest).Returns(txHash *common.Hash)`
| `eth_cancelPrivateTransaction` | `flashbots.CancelPrivateTx(txHash common.Hash).Returns(success *bool)`
| `flashbots_getUserStats`       | `flashbots.UserStats(blockNumber *big.Int).Returns(resp *flashbots.UserStatsResponse)`
| `flashbots_getBundleStats`     | `flashbots.BundleStats(bundleHash common.Hash, blockNumber *big.Int).Returns(resp *flashbots.BundleStatsResponse)`
