package keeper

import (
	"github.com/evgeniy-scherbina/cosmosevm/x/precompiles/types"
)

var _ types.QueryServer = Keeper{}
