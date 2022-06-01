package flashbots

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/core"
)

type SendPrivateTransactionRequest struct {
	Transaction    *types.Transaction // Signed transaction to send.
	RawTransaction []byte             // Raw signed transaction to send.
	MaxBlockNumber *big.Int           // Max block number for which the tx should be included (Optional).
	Fast           bool               // Enable fast mode (Optional). See https://docs.flashbots.net/flashbots-protect/rpc/fast-mode
}

type sendPrivateTransactionRequest struct {
	RawTransaction hexutil.Bytes `json:"tx"`
	MaxBlockNumber *hexutil.Big  `json:"maxBlockNumber"`
	Preferences    struct {
		Fast bool `json:"fast"`
	} `json:"preferences"`
}

// MarshalJSON implements the json.Marshaler interface.
func (c SendPrivateTransactionRequest) MarshalJSON() ([]byte, error) {
	var enc sendPrivateTransactionRequest

	if c.Transaction != nil {
		rawTx, err := c.Transaction.MarshalBinary()
		if err != nil {
			return nil, err
		}
		enc.RawTransaction = rawTx
	} else {
		enc.RawTransaction = c.RawTransaction
	}
	enc.MaxBlockNumber = (*hexutil.Big)(c.MaxBlockNumber)
	enc.Preferences.Fast = c.Fast
	return json.Marshal(&enc)
}

// SendPrivateTransaction sends a private transaction to the Flashbots relay.
func SendPrivateTransaction(r *SendPrivateTransactionRequest) core.CallerFactory[common.Hash] {
	return &sendPrivateTransactionFactory{params: r}
}

type sendPrivateTransactionFactory struct {
	// args
	params *SendPrivateTransactionRequest

	// returns
	returns *common.Hash
}

func (f *sendPrivateTransactionFactory) Returns(txHash *common.Hash) core.Caller {
	f.returns = txHash
	return f
}

func (f *sendPrivateTransactionFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_sendPrivateTransaction",
		Args:   []any{f.params},
		Result: f.returns,
	}, nil
}

func (f *sendPrivateTransactionFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	return nil
}
