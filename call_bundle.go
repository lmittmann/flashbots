package flashbots

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/flashbots/internal"
	"github.com/lmittmann/w3/w3types"
)

type CallBundleRequest struct {
	Transactions     types.Transactions // List of signed transactions to simulate in a bundle.
	RawTransactions  [][]byte           // List of signed raw transactions to simulate in a bundle.
	BlockNumber      *big.Int           // Block number for which the bundle is valid.
	StateBlockNumber *big.Int           // Block number of state to use for simulation, "latest" if nil.
	Timestamp        *big.Int           // Timestamp of block used for simulation (Optional).
}

type callBundleRequest struct {
	RawTransactions  []hexutil.Bytes `json:"txs"`
	BlockNumber      *hexutil.Big    `json:"blockNumber"`
	StateBlockNumber string          `json:"stateBlockNumber"`
	Timestamp        *big.Int        `json:"timestamp"`
}

// MarshalJSON implements the [json.Marshaler].
func (c CallBundleRequest) MarshalJSON() ([]byte, error) {
	var enc callBundleRequest

	if len(c.Transactions) > 0 {
		enc.RawTransactions = make([]hexutil.Bytes, len(c.Transactions))
		for i, tx := range c.Transactions {
			rawTx, err := tx.MarshalBinary()
			if err != nil {
				return nil, err
			}
			enc.RawTransactions[i] = rawTx
		}
	} else {
		enc.RawTransactions = make([]hexutil.Bytes, len(c.RawTransactions))
		for i, rawTx := range c.RawTransactions {
			enc.RawTransactions[i] = rawTx
		}
	}
	enc.BlockNumber = (*hexutil.Big)(c.BlockNumber)
	enc.StateBlockNumber = toBlockNumberArg(c.StateBlockNumber)
	enc.Timestamp = c.Timestamp
	return json.Marshal(&enc)
}

type CallBundleResponse struct {
	BundleGasPrice    *big.Int
	BundleHash        common.Hash
	CoinbaseDiff      *big.Int
	EthSentToCoinbase *big.Int
	GasFees           *big.Int
	StateBlockNumber  *big.Int
	TotalGasUsed      uint64
	Results           []CallBundleResult
}

type callBundleResponse struct {
	BundleGasPrice    *internal.StrInt   `json:"bundleGasPrice"`
	BundleHash        *common.Hash       `json:"bundleHash"`
	CoinbaseDiff      *internal.StrInt   `json:"coinbaseDiff"`
	EthSentToCoinbase *internal.StrInt   `json:"ethSentToCoinbase"`
	GasFees           *internal.StrInt   `json:"gasFees"`
	StateBlockNumber  *big.Int           `json:"stateBlockNumber"`
	TotalGasUsed      *uint64            `json:"totalGasUsed"`
	Results           []callBundleResult `json:"results"`
}

type CallBundleResult struct {
	CoinbaseDiff      *big.Int
	EthSentToCoinbase *big.Int
	FromAddress       common.Address
	GasFees           *big.Int
	GasPrice          *big.Int
	GasUsed           uint64
	ToAddress         *common.Address
	TxHash            common.Hash
	Value             []byte // Output

	Error  error
	Revert string // Revert reason
}

type callBundleResult struct {
	CoinbaseDiff      *internal.StrInt `json:"coinbaseDiff"`
	EthSentToCoinbase *internal.StrInt `json:"ethSentToCoinbase"`
	FromAddress       *common.Address  `json:"fromAddress"`
	GasFees           *internal.StrInt `json:"gasFees"`
	GasPrice          *internal.StrInt `json:"gasPrice"`
	GasUsed           *uint64          `json:"gasUsed"`
	ToAddress         *common.Address  `json:"toAddress"`
	TxHash            *common.Hash     `json:"txHash"`
	Value             *hexutil.Bytes   `json:"value"`

	Error  *string `json:"error"`
	Revert *string `json:"revert"`
}

// UnmarshalJSON implements the [json.Unmarshaler].
func (c *CallBundleResponse) UnmarshalJSON(input []byte) error {
	var dec callBundleResponse
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}

	if dec.BundleGasPrice != nil {
		c.BundleGasPrice = (*big.Int)(dec.BundleGasPrice)
	}
	if dec.BundleHash != nil {
		c.BundleHash = *dec.BundleHash
	}
	if dec.CoinbaseDiff != nil {
		c.CoinbaseDiff = (*big.Int)(dec.CoinbaseDiff)
	}
	if dec.EthSentToCoinbase != nil {
		c.EthSentToCoinbase = (*big.Int)(dec.EthSentToCoinbase)
	}
	if dec.GasFees != nil {
		c.GasFees = (*big.Int)(dec.GasFees)
	}
	if dec.StateBlockNumber != nil {
		c.StateBlockNumber = dec.StateBlockNumber
	}
	if dec.TotalGasUsed != nil {
		c.TotalGasUsed = *dec.TotalGasUsed
	}
	if dec.Results != nil {
		c.Results = make([]CallBundleResult, len(dec.Results))
		for i, res := range dec.Results {
			if res.CoinbaseDiff != nil {
				c.Results[i].CoinbaseDiff = (*big.Int)(res.CoinbaseDiff)
			}
			if res.EthSentToCoinbase != nil {
				c.Results[i].EthSentToCoinbase = (*big.Int)(res.EthSentToCoinbase)
			}
			if res.FromAddress != nil {
				c.Results[i].FromAddress = *res.FromAddress
			}
			if res.GasFees != nil {
				c.Results[i].GasFees = (*big.Int)(res.GasFees)
			}
			if res.GasPrice != nil {
				c.Results[i].GasPrice = (*big.Int)(res.GasPrice)
			}
			if res.GasUsed != nil {
				c.Results[i].GasUsed = *res.GasUsed
			}
			if res.ToAddress != nil {
				c.Results[i].ToAddress = res.ToAddress
			}
			if res.TxHash != nil {
				c.Results[i].TxHash = *res.TxHash
			}
			if res.Value != nil {
				c.Results[i].Value = *res.Value
			}
			if res.Error != nil {
				c.Results[i].Error = errors.New(*res.Error)
			}
			if res.Revert != nil {
				c.Results[i].Revert = *res.Revert
			}
		}
	}
	return nil
}

// CallBundle simulates a bundle.
func CallBundle(r *CallBundleRequest) w3types.CallerFactory[CallBundleResponse] {
	return &callBundleFactory{param: r}
}

type callBundleFactory struct {
	// args
	param *CallBundleRequest

	// returns
	returns *CallBundleResponse
}

func (f *callBundleFactory) Returns(resp *CallBundleResponse) w3types.Caller {
	f.returns = resp
	return f
}

// CreateRequest implements the [w3types.RequestCreator].
func (f *callBundleFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_callBundle",
		Args:   []any{f.param},
		Result: f.returns,
	}, nil
}

// HandleResponse implements the [w3types.ResponseHandler].
func (f *callBundleFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	return nil
}

func toBlockNumberArg(blockNumber *big.Int) string {
	if blockNumber == nil || blockNumber.Sign() < 0 {
		return "latest"
	}
	return hexutil.EncodeBig(blockNumber)
}
