package hashkeydid_go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const metaDataUrl = "https://api.hashkey.id/did/api/nft/metadata/%s"

type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// GetMetadataImage returns the image url in metadata by tokenId
func GetMetadataImage(tokenId string) (string, error) {
	metadata, err := GetMetadata(tokenId)
	if err != nil {
		return "", err
	}
	return metadata.Image, nil
}

// GetMetadata returns the Metadata by tokenId
func GetMetadata(tokenId string) (*Metadata, error) {
	client := &http.Client{}
	res, err := client.Get(fmt.Sprintf(metaDataUrl, tokenId))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var metadata Metadata
	err = json.Unmarshal(body, &metadata)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}
