package precompile_test

import (
	"testing"

	keepertest "github.com/evgeniy-scherbina/cosmosevm/testutil/keeper"
	"github.com/evgeniy-scherbina/cosmosevm/testutil/nullify"
	"github.com/evgeniy-scherbina/cosmosevm/x/precompile"
	"github.com/evgeniy-scherbina/cosmosevm/x/precompile/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PrecompileKeeper(t)
	precompile.InitGenesis(ctx, *k, genesisState)
	got := precompile.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
