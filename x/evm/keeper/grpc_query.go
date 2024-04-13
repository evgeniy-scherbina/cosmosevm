package keeper

import (
	"github.com/evgeniy-scherbina/cosmosevm/x/evm/types"
)

var _ types.QueryServer = Keeper{}
