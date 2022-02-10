//go:generate gencodec -type SendBundleParam -field-override sendBundleMarshaling -out send_bundle.gen.go
package flashbots

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type SendBundleParam struct {
	Transactions      types.Transactions `json:"-"`                               // List of signed transactions to execute in a bundle.
	RawTransactions   [][]byte           `json:"txs"         gencodec:"required"` // List of signed raw transactions to execute in a bundle.
	BlockNumber       *big.Int           `json:"blockNumber" gencodec:"required"` // Block number for which the bundle is valid
	MinTimestamp      *big.Int           `json:"minTimestamp,omitempty"`          // Minimum Unix Timestamp for which the bundle is valid
	MaxTimestamp      *big.Int           `json:"maxTimestamp,omitempty"`          // Maximum Unix Timestamp for which the bundle is valid
	RevertingTxHashes []common.Hash      `json:"revertingTxHashes,omitempty"`     // List of tx hashes in bundle that are allowed to revert.
}

type sendBundleMarshaling struct {
	RawTransactions   []hexutil.Bytes
	BlockNumber       *hexutil.Big
	MinTimestamp      *hexutil.Big
	MaxTimestamp      *hexutil.Big
	RevertingTxHashes []common.Hash
}

type sendBundleResultMarshaling struct {
	BundleHash common.Hash `json:"bundleHash"`
}

// SendBundle sends a bundle to the network.
func SendBundle(param *SendBundleParam) *SendBundleFactory {
	return &SendBundleFactory{param: param}
}

type SendBundleFactory struct {
	// args
	param *SendBundleParam

	// returns
	result  sendBundleResultMarshaling
	returns *common.Hash
}

func (f *SendBundleFactory) Returns(hash *common.Hash) *SendBundleFactory {
	f.returns = hash
	return f
}

// CreateRequest implements the w3/core.RequestCreator interface.
func (f *SendBundleFactory) CreateRequest() (rpc.BatchElem, error) {
	if lenTx := len(f.param.Transactions); lenTx > 0 {
		f.param.RawTransactions = make([][]byte, lenTx)
		for i, tx := range f.param.Transactions {
			rawTx, err := tx.MarshalBinary()
			if err != nil {
				return rpc.BatchElem{}, err
			}
			f.param.RawTransactions[i] = rawTx
		}
	}

	rawJson, err := f.param.MarshalJSON()
	if err != nil {
		return rpc.BatchElem{}, err
	}

	return rpc.BatchElem{
		Method: "eth_sendBundle",
		Args:   []interface{}{json.RawMessage(rawJson)},
		Result: &f.result,
	}, nil
}

// HandleResponse implements the w3/core.ResponseHandler interface.
func (f *SendBundleFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	*f.returns = f.result.BundleHash
	return nil
}
