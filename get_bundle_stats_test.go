package flashbots

import (
	"math/big"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/rpctest"
)

func TestBundleStats(t *testing.T) {
	t.Parallel()

	srv := rpctest.NewFileServer(t, "testdata/get_bundle_stats.golden")
	defer srv.Close()

	client := w3.MustDial(srv.URL())
	defer client.Close()

	var (
		bundleStats     = new(BundleStatsResponse)
		wantBundleStats = &BundleStatsResponse{
			IsSimulated:    true,
			IsSentToMiners: true,
			IsHighPriority: true,
			SimulatedAt:    mustParse("2021-08-06T21:36:06.317Z"),
			SubmittedAt:    mustParse("2021-08-06T21:36:06.250Z"),
			SentToMinersAt: mustParse("2021-08-06T21:36:06.343Z"),
		}
	)

	if err := client.Call(
		BundleStats(w3.H("0x2228f5d8954ce31dc1601a8ba264dbd401bf1428388ce88238932815c5d6f23f"), big.NewInt(999_999_999)).Returns(bundleStats),
	); err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	if diff := cmp.Diff(wantBundleStats, bundleStats); diff != "" {
		t.Fatalf("(-want, +got)\n%s", diff)
	}
}

func mustParse(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		panic(err.Error())
	}
	return t
}
