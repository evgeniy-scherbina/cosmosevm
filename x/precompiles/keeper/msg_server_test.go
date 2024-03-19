package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/evgeniy-scherbina/cosmosevm/x/precompiles/types"
    "github.com/evgeniy-scherbina/cosmosevm/x/precompiles/keeper"
    keepertest "github.com/evgeniy-scherbina/cosmosevm/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PrecompilesKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
