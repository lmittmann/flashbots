package flashbots

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/flashbots/internal"
	"github.com/lmittmann/w3/w3types"
)

type bundleStatsRequest struct {
	BundleHash  common.Hash  `json:"bundleHash"`
	BlockNumber *hexutil.Big `json:"blockNumber"`
}

type BundleStatsResponse struct {
	IsSimulated    bool
	IsSentToMiners bool
	IsHighPriority bool
	SimulatedAt    time.Time
	SubmittedAt    time.Time
	SentToMinersAt time.Time
}

// BundleStats requests the bundles Flashbots relay stats. The given block
// number must be within 20 blocks of the current chain tip.
func BundleStats(bundleHash common.Hash, blockNumber *big.Int) w3types.CallerFactory[BundleStatsResponse] {
	return &bundleStatsFactory{bundleHash: bundleHash, blockNumber: blockNumber}
}

type bundleStatsFactory struct {
	// args
	bundleHash  common.Hash
	blockNumber *big.Int

	// returns
	returns *BundleStatsResponse
}

func (f *bundleStatsFactory) Returns(bundleStats *BundleStatsResponse) w3types.Caller {
	f.returns = bundleStats
	return f
}

func (f *bundleStatsFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "flashbots_getBundleStats",
		Args: []any{&bundleStatsRequest{
			BundleHash:  f.bundleHash,
			BlockNumber: (*hexutil.Big)(f.blockNumber),
		}},
		Result: f.returns,
	}, nil
}

func (f *bundleStatsFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	return nil
}

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
	IsHighPriority       *bool            `json:"is_high_priority"`
	AllTimeMinerPayments *internal.StrInt `json:"all_time_miner_payments"`
	AllTimeGasSimulated  *internal.StrInt `json:"all_time_gas_simulated"`
	Last7dMinerPayments  *internal.StrInt `json:"last_7d_miner_payments"`
	Last7dGasSimulated   *internal.StrInt `json:"last_7d_gas_simulated"`
	Last1dMinerPayments  *internal.StrInt `json:"last_1d_miner_payments"`
	Last1dGasSimulated   *internal.StrInt `json:"last_1d_gas_simulated"`
}

// UnmarshalJSON implements the [json.Unmarshaler].
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
func UserStats(blockNumber *big.Int) w3types.CallerFactory[UserStatsResponse] {
	return &userStatsFactory{blockNumber: blockNumber}
}

type userStatsFactory struct {
	// args
	blockNumber *big.Int

	// returns
	returns *UserStatsResponse
}

func (f *userStatsFactory) Returns(userStats *UserStatsResponse) w3types.Caller {
	f.returns = userStats
	return f
}

func (f *userStatsFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "flashbots_getUserStats",
		Args:   []any{hexutil.EncodeBig(f.blockNumber)},
		Result: f.returns,
	}, nil
}

func (f *userStatsFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	return nil
}
