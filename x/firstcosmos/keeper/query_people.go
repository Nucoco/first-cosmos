package keeper

import (
	"context"

	"first-cosmos/x/firstcosmos/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PeopleAll(ctx context.Context, req *types.QueryAllPeopleRequest) (*types.QueryAllPeopleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var peoples []types.People

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	peopleStore := prefix.NewStore(store, types.KeyPrefix(types.PeopleKey))

	pageRes, err := query.Paginate(peopleStore, req.Pagination, func(key []byte, value []byte) error {
		var people types.People
		if err := k.cdc.Unmarshal(value, &people); err != nil {
			return err
		}

		peoples = append(peoples, people)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPeopleResponse{People: peoples, Pagination: pageRes}, nil
}

func (k Keeper) People(ctx context.Context, req *types.QueryGetPeopleRequest) (*types.QueryGetPeopleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	people, found := k.GetPeople(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPeopleResponse{People: people}, nil
}
