package hashkeydid_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// GetMetadataName returns the name in metadata by tokenId
func GetMetadataName(tokenId string) (string, error) {
	metadata, err := GetMetadata(tokenId)
	if err != nil {
		return "", err
	}
	return metadata.Name, nil
}

// GetMetadataDescription returns the description in metadata by tokenId
func GetMetadataDescription(tokenId string) (string, error) {
	metadata, err := GetMetadata(tokenId)
	if err != nil {
		return "", err
	}
	return metadata.Description, nil
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
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.hashkey.id/did/api/nft/metadata/%s", tokenId), bytes.NewReader([]byte{}))
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
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
