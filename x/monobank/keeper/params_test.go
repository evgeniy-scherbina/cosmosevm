package keeper_test

import (
	"testing"

	testkeeper "github.com/evgeniy-scherbina/cosmosevm/testutil/keeper"
	"github.com/evgeniy-scherbina/cosmosevm/x/monobank/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MonobankKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
