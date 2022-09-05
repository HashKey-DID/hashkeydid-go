package hashkeydid_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/hashkey-did/hashkeydid-go/contracts/erc1155"
	"github.com/hashkey-did/hashkeydid-go/contracts/erc721"
)

// GetMetadataAvatarByDIDName returns the image url in metadata queried by did name
func (c *Core) GetMetadataAvatarByDIDName(opts *bind.CallOpts, didName string) (string, error) {
	didClaimed, err := c.did.DidClaimed(opts, didName)
	if err != nil {
		return "", err
	}
	if !didClaimed {
		return "", ErrDidNotClaimed
	}
	tokenId, err := c.did.Did2TokenId(opts, didName)
	if err != nil {
		return "", err
	}
	return GetMetadataImage(tokenId.String())
}

// GetMetadataAvatarByTokenId returns the image url in metadata queried by tokenId
func (c *Core) GetMetadataAvatarByTokenId(opts *bind.CallOpts, tokenId uint64) (string, error) {
	_, err := c.did.OwnerOf(opts, new(big.Int).SetUint64(tokenId))
	if err != nil {
		if strings.Contains(err.Error(), "invalid token ID") {
			return "", ErrTokenIdNotMinted
		}
		return "", err
	}
	return GetMetadataImage(fmt.Sprintf("%d", tokenId))
}

// GetAvatarByDIDName returns the image url in resolver text queried by did name
func (c *Core) GetAvatarByDIDName(opts *bind.CallOpts, didName string, chainList map[string]*ChainInfo) (string, error) {
	didClaimed, err := c.did.DidClaimed(opts, didName)
	if err != nil {
		return "", err
	}
	if !didClaimed {
		return "", ErrDidNotClaimed
	}
	tokenId, err := c.did.Did2TokenId(opts, didName)
	if err != nil {
		return "", err
	}
	avatarText, err := c.resolver.Text(opts, tokenId, "avatar")
	if err != nil {
		return "", err
	}
	if avatarText == "" {
		return "", ErrAvatarNotSet
	}
	return avatarFormatText2AvatarUrl(opts, avatarText, chainList)
}

// GetAvatarByTokenId returns the image url in resolver text queried by tokenId
func (c *Core) GetAvatarByTokenId(opts *bind.CallOpts, tokenId uint64, chainList map[string]*ChainInfo) (string, error) {
	_, err := c.did.OwnerOf(opts, new(big.Int).SetUint64(tokenId))
	if err != nil {
		if strings.Contains(err.Error(), "invalid token ID") {
			return "", ErrTokenIdNotMinted
		}
		return "", err
	}
	avatarText, err := c.resolver.Text(opts, new(big.Int).SetUint64(tokenId), "avatar")
	if err != nil {
		return "", err
	}
	if avatarText == "" {
		return "", ErrAvatarNotSet
	}
	return avatarFormatText2AvatarUrl(opts, avatarText, chainList)
}

// avatarFormatText2AvatarUrl convert
func avatarFormatText2AvatarUrl(opts *bind.CallOpts, formatText string, chainList map[string]*ChainInfo) (string, error) {
	texts := strings.Split(formatText, ":")
	if len(texts) < 2 {
		return "", ErrInvalidAvatarText
	}
	switch texts[0] {
	case "nft":
		// nft:chainId:type(721/1155):contractAddr:tokenId
		if len(texts) != 5 {
			return "", ErrInvalidAvatarText
		}
		if chainList == nil {
			chainList = ChainList
		}
		client, err := ethclient.Dial(chainList[texts[1]].RPCUrl)
		if err != nil {
			return "", err
		}
		var tokenURI string
		tokenId, ok := new(big.Int).SetString(texts[4], 10)
		if !ok {
			return "", ErrInvalidAvatarText
		}
		switch texts[2] {
		case "721":
			nft721, err := erc721.NewContract(common.HexToAddress(texts[3]), client)
			if err != nil {
				return "", err
			}
			tokenURI, err = nft721.TokenURI(opts, tokenId)
		case "1155":
			nft1155, err := erc1155.NewContract(common.HexToAddress(texts[3]), client)
			if err != nil {
				return "", err
			}
			tokenURI, err = nft1155.Uri(opts, tokenId)
		default:
			return "", ErrInvalidAvatarText
		}
		image := getImageFromTokenURI(tokenURI)
		if image == "" {
			return "", ErrInvalidTokenURI
		}
		return image, nil
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "ipfs":
		if len(texts) != 2 {
			return "", ErrInvalidAvatarText
		}
		return fmt.Sprintf("%s://%s", texts[0], texts[1]), nil
	default:
		return "", ErrInvalidAvatarText
	}
}

func getImageFromTokenURI(tokenURI string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(tokenURI), bytes.NewReader([]byte{}))
	if err != nil {
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	var metadata Metadata
	err = json.Unmarshal(body, &metadata)
	if err != nil {
		return ""
	}
	return metadata.Image
}
