package keeper

import (
	"github.com/evgeniy-scherbina/cosmosevm/x/privatbank/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
