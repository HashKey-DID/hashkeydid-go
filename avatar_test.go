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
	_, err := core.GetMetadataAvatarByTokenId(opts, big.NewInt(15921))
	assert.Equal(t, ErrTokenIdNotMinted, err)
	opts.BlockNumber = new(big.Int).SetUint64(40069812)
	avatarUrl, err := core.GetMetadataAvatarByTokenId(opts, big.NewInt(15921))
	assert.Nil(t, err)
	assert.Equal(t, "https://api.hashkey.id/did/api/file/avatar_0c512334-53bb-40d6-9ef5-dedcf8269d19.png", avatarUrl)
}

func TestCore_GetAvatarByTokenId(t *testing.T) {
	core := initTestCore()
	avatarUrl1, err := core.GetAvatarByDIDName(nil, "herro.key")
	assert.Nil(t, err)
	avatarUrl2, err := core.GetAvatarByTokenId(nil, big.NewInt(13756))
	assert.Nil(t, err)
	assert.Equal(t, avatarUrl1, avatarUrl2)
}

func TestCore_avatarFormatText2AvatarUrl(t *testing.T) {
	url := "nft:1:721:0xdb0867214f0a2e129fbc8b72f2898851e28fb09f:1333"
	avatarUrl, err := avatarFormatText2AvatarUrl(nil, url)
	assert.Nil(t, err)
	assert.Equal(t, avatarUrl, "http://ipfs.io/ipfs/QmbwVE72NF8iuPnSdVent6oend3WXG7reYkA2GaaaCTEZy/1333.png")
}

func Test_getImageFromTokenURI(t *testing.T) {
	var tokenUri = "https://arweave.net/i8bb1L1T82q8mu8fN6zU4V-u6F_aaxtG263JHQPLAJY/1"
	image := getImageFromTokenURI(tokenUri)
	assert.Equal(t, image, "https://arweave.net/kfU89BXdRGr324UAccjYzUSJKjdXejBx4HnyesOn9qM")
}
