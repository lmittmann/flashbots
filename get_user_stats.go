package flashbots

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type UserStatsResponse struct {
	IsHighPriority       bool     // True if the searcher has an high enough reputation to be in the high priority queue.
	AllTimeMinerPayments *big.Int // Total amount paid to miners over all time.
	AllTimeGasSimulated  *big.Int // Total amount of gas simulated across all bundles submitted to the relay.
	Last7dMinerPayments  *big.Int // Total amount paid to miners over the last 7 days.
	Last7dGasSimulated   *big.Int // Total amount of gas simulated across all bundles submitted to the relay in the last 7 days.
	Last1dMinerPayments  *big.Int // Total amount paid to miners over the last day.
	Last1dGasSimulated   *big.Int // Total amount of gas simulated across all bundles submitted to the relay in the last day.
}

type userStatsResponse struct {
	IsHighPriority       *bool      `json:"is_high_priority"`
	AllTimeMinerPayments *strBigint `json:"all_time_miner_payments"`
	AllTimeGasSimulated  *strBigint `json:"all_time_gas_simulated"`
	Last7dMinerPayments  *strBigint `json:"last_7d_miner_payments"`
	Last7dGasSimulated   *strBigint `json:"last_7d_gas_simulated"`
	Last1dMinerPayments  *strBigint `json:"last_1d_miner_payments"`
	Last1dGasSimulated   *strBigint `json:"last_1d_gas_simulated"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *UserStatsResponse) UnmarshalJSON(input []byte) error {
	var dec userStatsResponse
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}

	if dec.IsHighPriority != nil {
		u.IsHighPriority = *dec.IsHighPriority
	}
	if dec.AllTimeMinerPayments != nil {
		u.AllTimeMinerPayments = (*big.Int)(dec.AllTimeMinerPayments)
	}
	if dec.AllTimeGasSimulated != nil {
		u.AllTimeGasSimulated = (*big.Int)(dec.AllTimeGasSimulated)
	}
	if dec.Last7dMinerPayments != nil {
		u.Last7dMinerPayments = (*big.Int)(dec.Last7dMinerPayments)
	}
	if dec.Last7dGasSimulated != nil {
		u.Last7dGasSimulated = (*big.Int)(dec.Last7dGasSimulated)
	}
	if dec.Last1dMinerPayments != nil {
		u.Last1dMinerPayments = (*big.Int)(dec.Last1dMinerPayments)
	}
	if dec.Last1dGasSimulated != nil {
		u.Last1dGasSimulated = (*big.Int)(dec.Last1dGasSimulated)
	}
	return nil
}

// UserStats requests the users Flashbots relay stats. The given block number
// must be within 20 blocks of the current chain tip.
func UserStats(blockNumber *big.Int) *UserStatsFactory {
	return &UserStatsFactory{blockNumber: blockNumber}
}

type UserStatsFactory struct {
	// args
	blockNumber *big.Int

	// returns
	result  UserStatsResponse
	returns *UserStatsResponse
}

func (f *UserStatsFactory) Returns(userStats *UserStatsResponse) *UserStatsFactory {
	f.returns = userStats
	return f
}

// CreateRequest implements the w3/core.RequestCreator interface.
func (f *UserStatsFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "flashbots_getUserStats",
		Args:   []interface{}{hexutil.EncodeBig(f.blockNumber)},
		Result: &f.result,
	}, nil
}

// HandleResponse implements the w3/core.ResponseHandler interface.
func (f *UserStatsFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	*f.returns = f.result
	return nil
}
