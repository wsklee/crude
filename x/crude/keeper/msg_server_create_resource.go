package keeper

import (
    "context"

    sdk "github.com/cosmos/cosmos-sdk/types"

    "crude/x/crude/types"
)

func (k msgServer) CreateResource(goCtx context.Context, msg *types.MsgCreateResource) (*types.MsgCreateResourceResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    var resource = types.Resource{
        Creator:  msg.Creator,
        Title:    msg.Title,
        Body:     msg.Body,
        Category: msg.Category,
    }
    id := k.AppendResource(
        ctx,
        resource,
    )
    return &types.MsgCreateResourceResponse{
        Id: id,
    }, nil
}
