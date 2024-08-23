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
	rpctest.RunTestCases(t, []rpctest.TestCase[*flashbots.BundleStatsResponse]{
		{
			Golden: "get_bundle_stats",
			Call:   flashbots.BundleStats(w3.H("0x2228f5d8954ce31dc1601a8ba264dbd401bf1428388ce88238932815c5d6f23f"), big.NewInt(999_999_999)),
			WantRet: &flashbots.BundleStatsResponse{
				IsSimulated:    true,
				IsSentToMiners: true,
				IsHighPriority: true,
				SimulatedAt:    mustParseTime("2021-08-06T21:36:06.317Z"),
				SubmittedAt:    mustParseTime("2021-08-06T21:36:06.250Z"),
				SentToMinersAt: mustParseTime("2021-08-06T21:36:06.343Z"),
			},
		},
	})
}

func TestBundleStatsV2(t *testing.T) {
	rpctest.RunTestCases(t, []rpctest.TestCase[*flashbots.BundleStatsV2Response]{
		{
			Golden: "get_bundle_stats_v2",
			Call:   flashbots.BundleStatsV2(w3.H("0x2228f5d8954ce31dc1601a8ba264dbd401bf1428388ce88238932815c5d6f23f"), big.NewInt(999_999_999)),
			WantRet: &flashbots.BundleStatsV2Response{
				IsHighPriority: true,
				IsSimulated:    true,
				SimulatedAt:    mustParseTime("2022-10-06T21:36:06.317Z"),
				ReceivedAt:     mustParseTime("2022-10-06T21:36:06.250Z"),
				ConsideredByBuildersAt: []*struct {
					Pubkey    string
					Timestamp time.Time
				}{
					{
						Pubkey:    "0x81babeec8c9f2bb9c329fd8a3b176032fe0ab5f3b92a3f44d4575a231c7bd9c31d10b6328ef68ed1e8c02a3dbc8e80f9",
						Timestamp: mustParseTime("2022-10-06T21:36:06.343Z"),
					},
					{
						Pubkey:    "0x81beef03aafd3dd33ffd7deb337407142c80fea2690e5b3190cfc01bde5753f28982a7857c96172a75a234cb7bcb994f",
						Timestamp: mustParseTime("2022-10-06T21:36:06.394Z"),
					},
					{
						Pubkey:    "0xa1dead1e65f0a0eee7b5170223f20c8f0cbf122eac3324d61afbdb33a8885ff8cab2ef514ac2c7698ae0d6289ef27fc",
						Timestamp: mustParseTime("2022-10-06T21:36:06.322Z"),
					},
				},
				SealedByBuildersAt: []*struct {
					Pubkey    string
					Timestamp time.Time
				}{
					{
						Pubkey:    "0x81beef03aafd3dd33ffd7deb337407142c80fea2690e5b3190cfc01bde5753f28982a7857c96172a75a234cb7bcb994f",
						Timestamp: mustParseTime("2022-10-06T21:36:07.742Z"),
					},
				},
			},
		},
	})
}

func TestUserStats(t *testing.T) {
	rpctest.RunTestCases(t, []rpctest.TestCase[*flashbots.UserStatsResponse]{
		{
			Golden: "get_user_stats",
			Call:   flashbots.UserStats(big.NewInt(999_999_999)),
			WantRet: &flashbots.UserStatsResponse{
				IsHighPriority:       true,
				AllTimeMinerPayments: w3.I("1280749594841588639"),
				AllTimeGasSimulated:  w3.I("30049470846"),
				Last7dMinerPayments:  w3.I("1280749594841588639"),
				Last7dGasSimulated:   w3.I("30049470846"),
				Last1dMinerPayments:  w3.I("142305510537954293"),
				Last1dGasSimulated:   w3.I("2731770076"),
			},
		},
	})
}

func TestUserStatsV2(t *testing.T) {
	rpctest.RunTestCases(t, []rpctest.TestCase[*flashbots.UserStatsV2Response]{
		{
			Golden: "get_user_stats_v2",
			Call:   flashbots.UserStatsV2(big.NewInt(999_999_999)),
			WantRet: &flashbots.UserStatsV2Response{
				IsHighPriority:           true,
				AllTimeValidatorPayments: w3.I("1280749594841588639"),
				AllTimeGasSimulated:      w3.I("30049470846"),
				Last7dValidatorPayments:  w3.I("1280749594841588639"),
				Last7dGasSimulated:       w3.I("30049470846"),
				Last1dValidatorPayments:  w3.I("142305510537954293"),
				Last1dGasSimulated:       w3.I("2731770076"),
			},
		},
	})
}

func mustParseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err.Error())
	}
	return t
}
