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

type SendPrivateTxRequest struct {
	Tx             *types.Transaction // Signed transaction to send.
	RawTx          []byte             // Raw signed transaction to send.
	MaxBlockNumber *big.Int           // Max block number for which the tx should be included (Optional).
	Fast           bool               // Enable fast mode (Optional). See https://docs.flashbots.net/flashbots-protect/rpc/fast-mode
}

type sendPrivateTxRequest struct {
	RawTx          hexutil.Bytes `json:"tx"`
	MaxBlockNumber *hexutil.Big  `json:"maxBlockNumber"`
	Preferences    struct {
		Fast bool `json:"fast"`
	} `json:"preferences"`
}

// MarshalJSON implements the json.Marshaler interface.
func (c SendPrivateTxRequest) MarshalJSON() ([]byte, error) {
	var enc sendPrivateTxRequest

	if c.Tx != nil {
		rawTx, err := c.Tx.MarshalBinary()
		if err != nil {
			return nil, err
		}
		enc.RawTx = rawTx
	} else {
		enc.RawTx = c.RawTx
	}
	enc.MaxBlockNumber = (*hexutil.Big)(c.MaxBlockNumber)
	enc.Preferences.Fast = c.Fast
	return json.Marshal(&enc)
}

// SendPrivateTx sends a private transaction to the Flashbots relay.
func SendPrivateTx(r *SendPrivateTxRequest) w3types.CallerFactory[common.Hash] {
	return &sendPrivateTxFactory{params: r}
}

type sendPrivateTxFactory struct {
	// args
	params *SendPrivateTxRequest

	// returns
	returns *common.Hash
}

func (f *sendPrivateTxFactory) Returns(txHash *common.Hash) w3types.Caller {
	f.returns = txHash
	return f
}

func (f *sendPrivateTxFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_sendPrivateTransaction",
		Args:   []any{f.params},
		Result: f.returns,
	}, nil
}

func (f *sendPrivateTxFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	return nil
}

type cancelPrivateTxRequest struct {
	TxHash common.Hash `json:"txHash"`
}

// CancelPrivateTx stops the private transactions with the given hash
// from being submitted for future blocks by the Flashbots relay.
func CancelPrivateTx(hash common.Hash) w3types.CallerFactory[bool] {
	return &cancelPrivateTxFactory{hash: hash}
}

type cancelPrivateTxFactory struct {
	// args
	hash common.Hash

	// returns
	returns *bool
}

func (f *cancelPrivateTxFactory) Returns(success *bool) w3types.Caller {
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
