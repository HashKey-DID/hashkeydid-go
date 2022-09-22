package hashkeydid_go

import "fmt"

var (
	ErrAddrNotClaimed    = fmt.Errorf("this address has not claimed a did")
	ErrAddrNotSetReverse = fmt.Errorf("this address has not set reverse record")
	ErrDidNotClaimed     = fmt.Errorf("this did name has not been claimed")
	ErrTokenIdNotMinted  = fmt.Errorf("this tokenId has not been minted")
	ErrAvatarNotSet      = fmt.Errorf("the avatar text has not been set on this did")
	ErrInvalidAvatarText = fmt.Errorf("the avatar text is invalid")
	ErrInvalidTokenURI   = fmt.Errorf("the tokenURI of avatar is invalid")
	ErrInvalidTokenId    = fmt.Errorf("the token id is invalid")
	ErrTimeExpired       = fmt.Errorf("the time has expired")
	ErrInvalidEvidence   = fmt.Errorf("the evidence is invalid")
)
