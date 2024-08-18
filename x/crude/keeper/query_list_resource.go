package keeper

import (
    "context"

    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    "github.com/cosmos/cosmos-sdk/types/query"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "crude/x/crude/types"
)

func (k Keeper) ListResource(ctx context.Context, req *types.QueryListResourceRequest) (*types.QueryListResourceResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))

    var resources []types.Resource
    pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
        var resource types.Resource
        if err := k.cdc.Unmarshal(value, &resource); err != nil {
            return err
        }

        // Filter by category if specified
        if req.Category == 0 || resource.Category == req.Category {
            resources = append(resources, resource)
        }
        return nil
    })

    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

    return &types.QueryListResourceResponse{Resource: resources, Pagination: pageRes}, nil
}
