package keeper

import (
	"github.com/evgeniy-scherbina/cosmosevm/x/privatbank/types"
)

var _ types.QueryServer = Keeper{}
