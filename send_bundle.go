package flashbots

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/w3types"
)

type SendBundleRequest struct {
	Transactions      types.Transactions // List of signed transactions to execute in a bundle.
	RawTransactions   [][]byte           // List of signed raw transactions to execute in a bundle.
	BlockNumber       *big.Int           // Block number for which the bundle is valid
	MinTimestamp      uint64             // Minimum Unix Timestamp for which the bundle is valid
	MaxTimestamp      uint64             // Maximum Unix Timestamp for which the bundle is valid
	RevertingTxHashes []common.Hash      // List of tx hashes in bundle that are allowed to revert.
	ReplacementUuid   string             // String, UUID that can be used to cancel/replace this bundle
}

type sendBundleRequest struct {
	RawTransactions   []hexutil.Bytes `json:"txs"`
	BlockNumber       *hexutil.Big    `json:"blockNumber"`
	MinTimestamp      uint64          `json:"minTimestamp,omitempty"`
	MaxTimestamp      uint64          `json:"maxTimestamp,omitempty"`
	RevertingTxHashes []common.Hash   `json:"revertingTxHashes,omitempty"`
	ReplacementUuid   string          `json:"replacementUuid,omitempty"`
}

// MarshalJSON implements the [json.Marshaler].
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
	if s.BlockNumber != nil {
		enc.BlockNumber = (*hexutil.Big)(s.BlockNumber)
	}
	enc.MinTimestamp = s.MinTimestamp
	enc.MaxTimestamp = s.MaxTimestamp
	enc.RevertingTxHashes = s.RevertingTxHashes
	enc.ReplacementUuid = s.ReplacementUuid
	return json.Marshal(&enc)
}

type sendBundleResponse struct {
	BundleHash common.Hash `json:"bundleHash"`
}

// SendBundle sends the bundle to the client's endpoint.
func SendBundle(r *SendBundleRequest) w3types.RPCCallerFactory[common.Hash] {
	return &sendBundleFactory{param: r}
}

type sendBundleFactory struct {
	// args
	param *SendBundleRequest

	// returns
	result  sendBundleResponse
	returns *common.Hash
}

func (f *sendBundleFactory) Returns(hash *common.Hash) w3types.RPCCaller {
	f.returns = hash
	return f
}

// CreateRequest implements the [w3types.RequestCreator].
func (f *sendBundleFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_sendBundle",
		Args:   []any{f.param},
		Result: &f.result,
	}, nil
}

// HandleResponse implements the [w3types.ResponseHandler].
func (f *sendBundleFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	if f.returns != nil {
		*f.returns = f.result.BundleHash
	}
	return nil
}
