package hashkeydid_go

import (
	"log"
)

const DefaultPlatONUrl = "https://openapi2.platon.network/rpc"

func initTestCore() *Core {
	core, err := NewDIDCore(DefaultPlatONUrl)
	if err != nil {
		log.Fatalf("init test core err: %s\n", err)
	}
	return core
}
