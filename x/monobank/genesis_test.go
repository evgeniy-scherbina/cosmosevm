package monobank_test

import (
	"testing"

	keepertest "github.com/evgeniy-scherbina/cosmosevm/testutil/keeper"
	"github.com/evgeniy-scherbina/cosmosevm/testutil/nullify"
	"github.com/evgeniy-scherbina/cosmosevm/x/monobank"
	"github.com/evgeniy-scherbina/cosmosevm/x/monobank/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		BalanceList: []types.Balance{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MonobankKeeper(t)
	monobank.InitGenesis(ctx, *k, genesisState)
	got := monobank.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.BalanceList, got.BalanceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
