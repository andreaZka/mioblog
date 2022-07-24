package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PostKeyPrefix is the prefix to retrieve all Post
	PostKeyPrefix = "Post/value/"
)

// PostKey returns the store key to retrieve a Post from the index fields
func PostKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
