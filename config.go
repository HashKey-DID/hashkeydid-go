package hashkeydid_go

type ChainInfo struct {
	ChainName       string
	ChainId         uint64
	DIDContract     string
	ResolveContract string
}

// ChainList is a map contains the infos of chain by the index of chainId
var ChainList = map[string]*ChainInfo{
	"1": {
		ChainName:       "Ethereum",
		ChainId:         1,
		DIDContract:     "0x7fdd3f96cbde51737a9e24b461e7e92a057c3bbf",
		ResolveContract: "0x2d64eb6bb9087a34281431eb0ff66a34ff7fa319",
	},
	"137": {
		ChainName:       "Polygon",
		ChainId:         137,
		DIDContract:     "0x7fdd3f96cbde51737a9e24b461e7e92a057c3bbf",
		ResolveContract: "0x2d64eb6bb9087a34281431eb0ff66a34ff7fa319",
	},
	"1284": {
		ChainName:       "Moonbeam",
		ChainId:         1284,
		DIDContract:     "0x39CE332189A4BF4D024528d6f7C90C5608B92c51",
		ResolveContract: "0xDad365540eAa3BE163F91aA5E32235cb00b06c73",
	},
	"210425": {
		ChainName:       "PlatON",
		ChainId:         210425,
		DIDContract:     "0x7fDd3f96cBDE51737A9E24b461E7E92A057C3BBf",
		ResolveContract: "0x606729294604A1c71f4BFc001894E4f8095Ec2eF",
	},
}
