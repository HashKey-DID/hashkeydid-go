package hashkeydid_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetMetadata(t *testing.T) {
	dstMetadata := Metadata{
		Name:        "gosdktest.key",
		Description: "Your Passport in the Metaverse",
		Image:       "https://api.hashkey.id/did/api/file/avatar_0c512334-53bb-40d6-9ef5-dedcf8269d19.png",
	}
	metadata, err := GetMetadata("15921")
	assert.Nil(t, err)
	assert.Equal(t, dstMetadata, *metadata)
}
