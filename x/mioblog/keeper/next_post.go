package keeper

import (
	"github.com/andreaZka/mioblog/x/mioblog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNextPost set nextPost in the store
func (k Keeper) SetNextPost(ctx sdk.Context, nextPost types.NextPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextPostKey))
	b := k.cdc.MustMarshal(&nextPost)
	store.Set([]byte{0}, b)
}

// GetNextPost returns nextPost
func (k Keeper) GetNextPost(ctx sdk.Context) (val types.NextPost, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextPostKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNextPost removes nextPost from the store
func (k Keeper) RemoveNextPost(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextPostKey))
	store.Delete([]byte{0})
}
