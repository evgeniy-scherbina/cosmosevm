package keeper

import (
	"github.com/evgeniy-scherbina/cosmosevm/x/precompile/types"
)

var _ types.QueryServer = Keeper{}
