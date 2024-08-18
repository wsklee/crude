package keeper

import (
    "context"
    "fmt"

    errorsmod "cosmossdk.io/errors"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

    "crude/x/crude/types"
)

func (k msgServer) DeleteResource(goCtx context.Context, msg *types.MsgDeleteResource) (*types.MsgDeleteResourceResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    val, found := k.GetResource(ctx, msg.Id)
    if !found {
        return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }
    if msg.Creator != val.Creator {
        return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }
    k.RemoveResource(ctx, msg.Id)
    return &types.MsgDeleteResourceResponse{}, nil
}