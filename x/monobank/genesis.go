package monobank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evgeniy-scherbina/cosmosevm/x/monobank/keeper"
	"github.com/evgeniy-scherbina/cosmosevm/x/monobank/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the balance
	for _, elem := range genState.BalanceList {
		k.SetBalance(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.BalanceList = k.GetAllBalance(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
