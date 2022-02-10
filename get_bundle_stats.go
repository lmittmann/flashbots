package flashbots

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type bundleStatMarshaling struct {
	BundleHash  common.Hash  `json:"bundleHash"`
	BlockNumber *hexutil.Big `json:"blockNumber"`
}

type BundleStatsResponse struct {
	IsSimulated    bool      `json:"isSimulated"`
	IsSentToMiners bool      `json:"isSentToMiners"`
	IsHighPriority bool      `json:"isHighPriority"`
	SimulatedAt    time.Time `json:"simulatedAt"`
	SubmittedAt    time.Time `json:"submittedAt"`
	SentToMinersAt time.Time `json:"sentToMinersAt"`
}

// BundleStats requests the bundles Flashbots relay stats. The given block
// number must be within 20 blocks of the current chain tip.
func BundleStats(bundleHash common.Hash, blockNumber *big.Int) *BundleStatsFactory {
	return &BundleStatsFactory{bundleHash: bundleHash, blockNumber: blockNumber}
}

type BundleStatsFactory struct {
	// args
	bundleHash  common.Hash
	blockNumber *big.Int

	// returns
	result  BundleStatsResponse
	returns *BundleStatsResponse
}

func (f *BundleStatsFactory) Returns(bundleStats *BundleStatsResponse) *BundleStatsFactory {
	f.returns = bundleStats
	return f
}

// CreateRequest implements the w3/core.RequestCreator interface.
func (f *BundleStatsFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "flashbots_getBundleStats",
		Args: []interface{}{&bundleStatMarshaling{
			BundleHash:  f.bundleHash,
			BlockNumber: (*hexutil.Big)(f.blockNumber),
		}},
		Result: &f.result,
	}, nil
}

// HandleResponse implements the w3/core.ResponseHandler interface.
func (f *BundleStatsFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	*f.returns = f.result
	return nil
}
