package keeper

import (
	"context"

	"first-cosmos/x/firstcosmos/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Hello(goCtx context.Context, msg *types.MsgHello) (*types.MsgHelloResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgHelloResponse{}, nil
}
