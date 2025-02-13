package balancer_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/osmosis-labs/osmosis/v7/x/gamm/pool-models/balancer"
	"github.com/osmosis-labs/osmosis/v7/x/gamm/types"
)

func createTestPool(t *testing.T, poolAssets []balancer.PoolAsset, swapFee, exitFee string) types.PoolI {
	swapFeeDec, err := sdk.NewDecFromStr(swapFee)
	require.NoError(t, err)

	exitFeeDec, err := sdk.NewDecFromStr(exitFee)
	require.NoError(t, err)

	pool, err := balancer.NewBalancerPool(1, balancer.PoolParams{
		SwapFee: swapFeeDec,
		ExitFee: exitFeeDec,
	}, poolAssets, "", time.Now())

	require.NoError(t, err)

	return &pool
}

func createTestContext(t *testing.T) sdk.Context {
	db := dbm.NewMemDB()
	logger := log.NewNopLogger()

	ms := rootmulti.NewStore(db, logger)

	return sdk.NewContext(ms, tmtypes.Header{}, false, logger)
}
