package keeper

import (
	"github.com/andreaZka/mioblog/x/mioblog/types"
)

var _ types.QueryServer = Keeper{}
