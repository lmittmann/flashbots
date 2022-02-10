package flashbots

import (
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/rpctest"
)

func TestUserStats(t *testing.T) {
	t.Parallel()

	srv := rpctest.NewFileServer(t, "testdata/get_user_stats.golden")
	defer srv.Close()

	client := w3.MustDial(srv.URL())
	defer client.Close()

	var (
		userStats     = new(UserStatsResponse)
		wantUserStats = &UserStatsResponse{
			IsHighPriority:       true,
			AllTimeMinerPayments: w3.I("1280749594841588639"),
			AllTimeGasSimulated:  w3.I("30049470846"),
			Last7dMinerPayments:  w3.I("1280749594841588639"),
			Last7dGasSimulated:   w3.I("30049470846"),
			Last1dMinerPayments:  w3.I("142305510537954293"),
			Last1dGasSimulated:   w3.I("2731770076"),
		}
	)

	if err := client.Call(UserStats(big.NewInt(999_999_999)).Returns(userStats)); err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	if diff := cmp.Diff(wantUserStats, userStats,
		cmp.AllowUnexported(big.Int{}),
	); diff != "" {
		t.Fatalf("(-want, +got)\n%s", diff)
	}
}
