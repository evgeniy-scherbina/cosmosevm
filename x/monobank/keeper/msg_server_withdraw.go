package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/evgeniy-scherbina/cosmosevm/x/monobank/testutil"
	"github.com/evgeniy-scherbina/cosmosevm/x/monobank/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	withdrawer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	balance, found := k.GetBalance(ctx, withdrawer.String())
	if !found {
		balance = types.Balance{
			Address: withdrawer.String(),
			Balance: 0,
		}
	}
	if balance.Balance < msg.Amount {
		return nil, fmt.Errorf("insufficient funds")
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, withdrawer, testutil.CoinsOf(msg.Amount))
	if err != nil {
		return nil, err
	}

	balance.Balance -= msg.Amount
	k.SetBalance(ctx, balance)

	return &types.MsgWithdrawResponse{}, nil
}
