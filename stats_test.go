package flashbots_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/lmittmann/flashbots"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/rpctest"
)

func TestBundleStats(t *testing.T) {
	tests := []rpctest.TestCase[flashbots.BundleStatsResponse]{
		{
			Golden: "get_bundle_stats",
			Call:   flashbots.BundleStats(w3.H("0x2228f5d8954ce31dc1601a8ba264dbd401bf1428388ce88238932815c5d6f23f"), big.NewInt(999_999_999)),
			WantRet: flashbots.BundleStatsResponse{
				IsSimulated:    true,
				IsSentToMiners: true,
				IsHighPriority: true,
				SimulatedAt:    mustParse("2021-08-06T21:36:06.317Z"),
				SubmittedAt:    mustParse("2021-08-06T21:36:06.250Z"),
				SentToMinersAt: mustParse("2021-08-06T21:36:06.343Z"),
			},
		},
	}

	rpctest.RunTestCases(t, tests)
}

func mustParse(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		panic(err.Error())
	}
	return t
}

func TestUserStats(t *testing.T) {
	tests := []rpctest.TestCase[flashbots.UserStatsResponse]{
		{
			Golden: "get_user_stats",
			Call:   flashbots.UserStats(big.NewInt(999_999_999)),
			WantRet: flashbots.UserStatsResponse{
				IsHighPriority:       true,
				AllTimeMinerPayments: w3.I("1280749594841588639"),
				AllTimeGasSimulated:  w3.I("30049470846"),
				Last7dMinerPayments:  w3.I("1280749594841588639"),
				Last7dGasSimulated:   w3.I("30049470846"),
				Last1dMinerPayments:  w3.I("142305510537954293"),
				Last1dGasSimulated:   w3.I("2731770076"),
			},
		},
	}

	rpctest.RunTestCases(t, tests)
}
