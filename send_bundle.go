package flashbots

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type SendBundleRequest struct {
	Transactions      types.Transactions // List of signed transactions to execute in a bundle.
	RawTransactions   [][]byte           // List of signed raw transactions to execute in a bundle.
	BlockNumber       *big.Int           // Block number for which the bundle is valid
	MinTimestamp      *big.Int           // Minimum Unix Timestamp for which the bundle is valid
	MaxTimestamp      *big.Int           // Maximum Unix Timestamp for which the bundle is valid
	RevertingTxHashes []common.Hash      // List of tx hashes in bundle that are allowed to revert.
}

type sendBundleRequest struct {
	RawTransactions   []hexutil.Bytes `json:"txs"`
	BlockNumber       *hexutil.Big    `json:"blockNumber"`
	MinTimestamp      *hexutil.Big    `json:"minTimestamp,omitempty"`
	MaxTimestamp      *hexutil.Big    `json:"maxTimestamp,omitempty"`
	RevertingTxHashes []common.Hash   `json:"revertingTxHashes,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface.
func (s SendBundleRequest) MarshalJSON() ([]byte, error) {
	var enc sendBundleRequest

	if len(s.Transactions) > 0 {
		enc.RawTransactions = make([]hexutil.Bytes, len(s.Transactions))
		for i, tx := range s.Transactions {
			rawTx, err := tx.MarshalBinary()
			if err != nil {
				return nil, err
			}
			enc.RawTransactions[i] = rawTx
		}
	} else {
		enc.RawTransactions = make([]hexutil.Bytes, len(s.RawTransactions))
		for i, rawTx := range s.RawTransactions {
			enc.RawTransactions[i] = rawTx
		}
	}
	enc.BlockNumber = (*hexutil.Big)(s.BlockNumber)
	enc.MinTimestamp = (*hexutil.Big)(s.MinTimestamp)
	enc.MaxTimestamp = (*hexutil.Big)(s.MaxTimestamp)
	enc.RevertingTxHashes = s.RevertingTxHashes
	return json.Marshal(&enc)
}

type sendBundleResponse struct {
	BundleHash common.Hash `json:"bundleHash"`
}

// SendBundle sends the bundle to the client's endpoint.
func SendBundle(r *SendBundleRequest) *SendBundleFactory {
	return &SendBundleFactory{param: r}
}

type SendBundleFactory struct {
	// args
	param *SendBundleRequest

	// returns
	result  sendBundleResponse
	returns *common.Hash
}

func (f *SendBundleFactory) Returns(hash *common.Hash) *SendBundleFactory {
	f.returns = hash
	return f
}

// CreateRequest implements the w3/core.RequestCreator interface.
func (f *SendBundleFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_sendBundle",
		Args:   []interface{}{f.param},
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
