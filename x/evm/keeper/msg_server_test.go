package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/evgeniy-scherbina/cosmosevm/testutil/keeper"
	"github.com/evgeniy-scherbina/cosmosevm/x/evm/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EvmKeeper(t)
	return k, sdk.WrapSDKContext(ctx)
}
