package hashkeydid_go

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/hashkeydid/hashkeydid-go/contracts/erc1155"
	"github.com/hashkeydid/hashkeydid-go/contracts/erc721"
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
func (c *Core) GetMetadataAvatarByTokenId(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	_, err := c.did.OwnerOf(opts, tokenId)
	if err != nil {
		if strings.Contains(err.Error(), "invalid token ID") {
			return "", ErrTokenIdNotMinted
		}
		return "", err
	}
	return GetMetadataImage(fmt.Sprintf("%d", tokenId))
}

// GetAvatarByDIDName returns the image url in resolver text queried by did name
func (c *Core) GetAvatarByDIDName(opts *bind.CallOpts, didName string, chainRpc ...*ChainRpcInfo) (string, error) {
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
	return avatarFormatText2AvatarUrl(opts, avatarText, chainRpc...)
}

// GetAvatarByTokenId returns the image url in resolver text queried by tokenId
func (c *Core) GetAvatarByTokenId(opts *bind.CallOpts, tokenId *big.Int, chainRpc ...*ChainRpcInfo) (string, error) {
	_, err := c.did.OwnerOf(opts, tokenId)
	if err != nil {
		if strings.Contains(err.Error(), "invalid token ID") {
			return "", ErrTokenIdNotMinted
		}
		return "", err
	}
	avatarText, err := c.resolver.Text(opts, tokenId, "avatar")
	if err != nil {
		return "", err
	}
	if avatarText == "" {
		return "", ErrAvatarNotSet
	}
	return avatarFormatText2AvatarUrl(opts, avatarText, chainRpc...)
}

// avatarFormatText2AvatarUrl convert avatar format text in resolver to an image url
func avatarFormatText2AvatarUrl(opts *bind.CallOpts, formatText string, chainRpc ...*ChainRpcInfo) (string, error) {
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

		chainList := make(map[string]*ChainRpcInfo)
		if len(chainRpc) == 0 {
			chainList = ChainRpcMap
		} else {
			for _, info := range chainRpc {
				chainList[strconv.FormatUint(info.ChainId, 10)] = info
			}
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
			if err != nil {
				return "", err
			}
		case "1155":
			nft1155, err := erc1155.NewContract(common.HexToAddress(texts[3]), client)
			if err != nil {
				return "", err
			}
			tokenURI, err = nft1155.Uri(opts, tokenId)
			if err != nil {
				return "", err
			}
		default:
			return "", ErrInvalidAvatarText
		}
		image := getImageFromTokenURI(tokenURI)
		if image == "" {
			return "", ErrInvalidTokenURI
		}
		return image, nil
	case "ipfs", "https", "http":
		if len(texts) != 2 {
			return "", ErrInvalidAvatarText
		}
		return fmt.Sprintf("%s:%s", texts[0], texts[1]), nil
	default:
		return "", ErrInvalidAvatarText
	}
}

// getImageFromTokenURI parses tokenURI's info to get the image url
func getImageFromTokenURI(tokenURI string) string {
	//client := http.Client{
	//	//Transport: &http.Transport{
	//	//	TLSClientConfig: &tls.Config{
	//	//		InsecureSkipVerify: true,
	//	//	},
	//	//},
	//}
	res, err := http.Get(tokenURI)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
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
