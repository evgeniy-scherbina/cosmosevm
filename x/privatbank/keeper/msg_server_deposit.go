package keeper

import (
	"context"
	"github.com/evgeniy-scherbina/cosmosevm/x/privatbank/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evgeniy-scherbina/cosmosevm/x/privatbank/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	depositor, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	err = k.bank.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, testutil.CoinsOf(msg.Amount))
	if err != nil {
		return nil, err
	}

	balance, found := k.GetBalance(ctx, depositor.String())
	if !found {
		balance = types.Balance{
			Address: depositor.String(),
			Balance: 0,
		}
	}

	balance.Balance += msg.Amount
	k.SetBalance(ctx, balance)

	return &types.MsgDepositResponse{}, nil
}
