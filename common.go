package flashbots

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func toBlockNumberArg(blockNumber *big.Int) string {
	if blockNumber == nil || blockNumber.Sign() < 0 {
		return "latest"
	}
	return hexutil.EncodeBig(blockNumber)
}
