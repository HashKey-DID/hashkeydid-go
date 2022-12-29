package hashkeydid_go

import (
	"fmt"
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

func TestCore_GetAvatarByTokenId(t *testing.T){
	core := initTestCore()
	avatarUrl1, err := core.GetAvatarByDIDName(nil, "gosdktest.key", nil)
	avatarUrl2, err := core.GetAvatarByTokenId(nil, 13756, nil)
	fmt.Println(avatarUrl1 == avatarUrl2)
	assert.Nil(t, err)
}

func TestCore_avatarFormatText2AvatarUrl(t *testing.T){
	// url := "nft:1:721:0xd2ff891f5556c623f36a3f22b0e4815a3e36dc23:1"
	url := "nft:1:721:0xdb0867214f0a2e129fbc8b72f2898851e28fb09f:1333"
	fmt.Println(avatarFormatText2AvatarUrl(nil, url, nil))
}

func Test_getImageFromTokenURI(t *testing.T){
	// var proxyUrl = "http://127.0.0.1:1080"
	var tokenUri = "https://arweave.net/i8bb1L1T82q8mu8fN6zU4V-u6F_aaxtG263JHQPLAJY/1"
	getImageFromTokenURI(tokenUri)
}
