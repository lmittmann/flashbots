package flashbots

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/core"
)

type cancelPrivateTransactionRequest struct {
	TxHash common.Hash `json:"txHash"`
}

// CancelPrivateTransaction stops the private transactions with the given hash
// from being submitted for future blocks by the Flashbots relay.
func CancelPrivateTransaction(hash common.Hash) core.CallerFactory[bool] {
	return &cancelPrivateTransactionFactory{hash: hash}
}

type cancelPrivateTransactionFactory struct {
	// args
	hash common.Hash

	// returns
	returns *bool
}

func (f *cancelPrivateTransactionFactory) Returns(success *bool) core.Caller {
	f.returns = success
	return f
}

func (f *cancelPrivateTransactionFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_cancelPrivateTransaction",
		Args: []any{&cancelPrivateTransactionRequest{
			TxHash: f.hash,
		}},
		Result: &f.returns,
	}, nil
}

func (f *cancelPrivateTransactionFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	return nil
}
