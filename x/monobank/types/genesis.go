package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		BalanceList: []Balance{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in balance
	balanceIndexMap := make(map[string]struct{})

	for _, elem := range gs.BalanceList {
		index := string(BalanceKey(elem.Address))
		if _, ok := balanceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for balance")
		}
		balanceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
