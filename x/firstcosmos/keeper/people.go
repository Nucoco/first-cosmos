package keeper

import (
	"context"
	"encoding/binary"

	"first-cosmos/x/firstcosmos/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetPeopleCount get the total number of people
func (k Keeper) GetPeopleCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PeopleCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPeopleCount set the total number of people
func (k Keeper) SetPeopleCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PeopleCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPeople appends a people in the store with a new id and update the count
func (k Keeper) AppendPeople(
	ctx context.Context,
	people types.People,
) uint64 {
	// Create the people
	count := k.GetPeopleCount(ctx)

	// Set the ID of the appended value
	people.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PeopleKey))
	appendedValue := k.cdc.MustMarshal(&people)
	store.Set(GetPeopleIDBytes(people.Id), appendedValue)

	// Update people count
	k.SetPeopleCount(ctx, count+1)

	return count
}

// SetPeople set a specific people in the store
func (k Keeper) SetPeople(ctx context.Context, people types.People) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PeopleKey))
	b := k.cdc.MustMarshal(&people)
	store.Set(GetPeopleIDBytes(people.Id), b)
}

// GetPeople returns a people from its id
func (k Keeper) GetPeople(ctx context.Context, id uint64) (val types.People, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PeopleKey))
	b := store.Get(GetPeopleIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePeople removes a people from the store
func (k Keeper) RemovePeople(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PeopleKey))
	store.Delete(GetPeopleIDBytes(id))
}

// GetAllPeople returns all people
func (k Keeper) GetAllPeople(ctx context.Context) (list []types.People) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PeopleKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.People
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPeopleIDBytes returns the byte representation of the ID
func GetPeopleIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.PeopleKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
