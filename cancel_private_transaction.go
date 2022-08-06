package flashbots

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/core"
)

type cancelPrivateTxRequest struct {
	TxHash common.Hash `json:"txHash"`
}

// CancelPrivateTx stops the private transactions with the given hash
// from being submitted for future blocks by the Flashbots relay.
func CancelPrivateTx(hash common.Hash) core.CallerFactory[bool] {
	return &cancelPrivateTxFactory{hash: hash}
}

type cancelPrivateTxFactory struct {
	// args
	hash common.Hash

	// returns
	returns *bool
}

func (f *cancelPrivateTxFactory) Returns(success *bool) core.Caller {
	f.returns = success
	return f
}

func (f *cancelPrivateTxFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_cancelPrivateTransaction",
		Args: []any{&cancelPrivateTxRequest{
			TxHash: f.hash,
		}},
		Result: &f.returns,
	}, nil
}

func (f *cancelPrivateTxFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	return nil
}
