package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PostList: []Post{},
		NextPost: nil,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in post
	postIndexMap := make(map[string]struct{})

	for _, elem := range gs.PostList {
		index := string(PostKey(elem.Index))
		if _, ok := postIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for post")
		}
		postIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
