package hashkeydid_go

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// GetMetadataAvatarByDIDName returns the image url in metadata searched by did name
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

// GetMetadataAvatarByTokenId returns the image url in metadata searched by tokenId
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

// GetAvatarByDIDName returns the image url in resolver text searched by did name
func (c *Core) GetAvatarByDIDName(opts *bind.CallOpts, didName string) (string, error) {
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
	return avatarText, nil
}

// GetAvatarByTokenId returns the image url in resolver text searched by tokenId
func (c *Core) GetAvatarByTokenId(opts *bind.CallOpts, tokenId uint64) (string, error) {
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
	return avatarText, nil
}

// avatarFormatText2AvatarUrl convert
func avatarFormatText2AvatarUrl(formatText string) (string, error) {
	texts := strings.Split(formatText, ":")
	if len(texts) < 2 {
		return "", ErrInvalidAvatarText
	}
	switch texts[0] {
	case "nft":
		if len(texts) != 4 {
			return "", ErrInvalidAvatarText
		}
		// todo wait to deal with formatText
		return formatText, nil
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
