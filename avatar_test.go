package hashkeydid_go

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
)

func TestCore_GetMetadataAvatarByDIDName(t *testing.T) {
	core := initTestCore()
	// the did is claimed at 40069812
	opts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(40069811)}
	_, err := core.GetMetadataAvatarByDIDName(opts, "gosdktest.key")
	assert.Equal(t, ErrDidNotClaimed, err)
	opts.BlockNumber = new(big.Int).SetUint64(40069812)
	avatarUrl, err := core.GetMetadataAvatarByDIDName(opts, "gosdktest.key")
	assert.Nil(t, err)
	assert.Equal(t, "https://api.hashkey.id/did/api/file/avatar_0c512334-53bb-40d6-9ef5-dedcf8269d19.png", avatarUrl)
}

func TestCore_GetMetadataAvatarByTokenId(t *testing.T) {
	core := initTestCore()
	// the did is claimed at 40069812
	opts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(40069811)}
	_, err := core.GetMetadataAvatarByTokenId(opts, 15921)
	assert.Equal(t, ErrTokenIdNotMinted, err)
	opts.BlockNumber = new(big.Int).SetUint64(40069812)
	avatarUrl, err := core.GetMetadataAvatarByTokenId(opts, 15921)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.hashkey.id/did/api/file/avatar_0c512334-53bb-40d6-9ef5-dedcf8269d19.png", avatarUrl)
}
