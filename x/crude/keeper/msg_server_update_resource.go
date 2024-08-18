package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"crude/x/crude/types"
)

func (k msgServer) UpdateResource(goCtx context.Context, msg *types.MsgUpdateResource) (*types.MsgUpdateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var resource = types.Resource{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Title:    msg.Title,
		Body:     msg.Body,
		Category: msg.Category,
		NewField: msg.NewField, // Include the new field
	}
	val, found := k.GetResource(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetResource(ctx, resource)
	return &types.MsgUpdateResourceResponse{}, nil
}
