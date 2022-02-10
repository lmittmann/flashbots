//go:generate gencodec -type UserStatsResponse -field-override userStatsResponseMarshaling -out get_user_stats.gen.go
package flashbots

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type UserStatsResponse struct {
	IsHighPriority       bool     `json:"is_high_priority"`
	AllTimeMinerPayments *big.Int `json:"all_time_miner_payments"`
	AllTimeGasSimulated  *big.Int `json:"all_time_gas_simulated"`
	Last7dMinerPayments  *big.Int `json:"last_7d_miner_payments"`
	Last7dGasSimulated   *big.Int `json:"last_7d_gas_simulated"`
	Last1dMinerPayments  *big.Int `json:"last_1d_miner_payments"`
	Last1dGasSimulated   *big.Int `json:"last_1d_gas_simulated"`
}

type userStatsResponseMarshaling struct {
	IsHighPriority       bool
	AllTimeMinerPayments *strBigint
	AllTimeGasSimulated  *strBigint
	Last7dMinerPayments  *strBigint
	Last7dGasSimulated   *strBigint
	Last1dMinerPayments  *strBigint
	Last1dGasSimulated   *strBigint
}

// UserStats requests the users Flashbots realy stats. The given block number
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
