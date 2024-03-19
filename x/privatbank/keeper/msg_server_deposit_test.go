package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/evgeniy-scherbina/cosmosevm/testutil/keeper"
	"github.com/evgeniy-scherbina/cosmosevm/x/privatbank/keeper"
	"github.com/evgeniy-scherbina/cosmosevm/x/privatbank/testutil"
	"github.com/evgeniy-scherbina/cosmosevm/x/privatbank/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDepositAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankEscrowKeeper(ctrl)

	k, ctx := keepertest.PrivatbankKeeperWithMocks(t, bankMock)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	aliceAddr, err := sdk.AccAddressFromBech32(testutil.Alice)
	require.NoError(t, err)

	bankMock.EXPECT().
		SendCoinsFromAccountToModule(ctx, aliceAddr, types.ModuleName, testutil.CoinsOf(1)).
		AnyTimes()
	bankMock.EXPECT().
		SendCoinsFromAccountToModule(ctx, aliceAddr, types.ModuleName, testutil.CoinsOf(2)).
		AnyTimes()

	//bankMock.EXPECT().
	//	SendCoinsFromModuleToAccount(ctx, types.ModuleName, aliceAddr, testutil.CoinsOf(1)).
	//	AnyTimes()
	//bankMock.EXPECT().
	//	SendCoinsFromModuleToAccount(ctx, types.ModuleName, aliceAddr, testutil.CoinsOf(2)).
	//	AnyTimes()

	// Deposit 1 token
	{
		_, err = msgServer.Deposit(goCtx, &types.MsgDeposit{
			Creator: testutil.Alice,
			Amount:  1,
		})
		require.NoError(t, err)

		balance, found := k.GetBalance(ctx, testutil.Alice)
		require.True(t, found)
		require.Equal(t, uint64(1), balance.Balance)
	}

	// Deposit 2 more tokens
	{
		_, err = msgServer.Deposit(goCtx, &types.MsgDeposit{
			Creator: testutil.Alice,
			Amount:  2,
		})
		require.NoError(t, err)

		balance, found := k.GetBalance(ctx, testutil.Alice)
		require.True(t, found)
		require.Equal(t, uint64(3), balance.Balance)
	}

	//// Withdraw 1 token
	//{
	//	_, err = msgServer.Withdraw(goCtx, &types.MsgWithdraw{
	//		Creator: testutil.Alice,
	//		Amount:  1,
	//	})
	//	require.NoError(t, err)
	//
	//	balance, found := k.GetBalance(ctx, testutil.Alice)
	//	require.True(t, found)
	//	require.Equal(t, uint64(2), balance.Balance)
	//}
	//
	//// Withdraw 2 more tokens
	//{
	//	_, err = msgServer.Withdraw(goCtx, &types.MsgWithdraw{
	//		Creator: testutil.Alice,
	//		Amount:  2,
	//	})
	//	require.NoError(t, err)
	//
	//	balance, found := k.GetBalance(ctx, testutil.Alice)
	//	require.True(t, found)
	//	require.Equal(t, uint64(0), balance.Balance)
	//}
}
