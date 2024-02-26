package privatbank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evgeniy-scherbina/cosmosevm/x/privatbank/keeper"
	"github.com/evgeniy-scherbina/cosmosevm/x/privatbank/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	return genesis
}