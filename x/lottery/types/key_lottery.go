package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// LotteryKeyPrefix is the prefix to retrieve all Lottery
	LotteryKeyPrefix = "Lottery/value/"
)

// LotteryKey returns the store key to retrieve a Lottery from the index fields
func LotteryKey(
	index uint64,
) []byte {
	var key []byte

	indexBytes := sdk.Uint64ToBigEndian(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
