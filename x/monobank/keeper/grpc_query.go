package keeper

import (
	"github.com/evgeniy-scherbina/cosmosevm/x/monobank/types"
)

var _ types.QueryServer = Keeper{}
