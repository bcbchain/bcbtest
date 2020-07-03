package resource

import (
	"github.com/bcbchain/sdk/sdk/types"
)

type Node struct {
	Name          string
	Url           string
	PubKey        types.PubKey
	RewardAddress types.Address
	Power         int64
}

var (
	genesisNodes  []Node
	observerNodes []Node
	freeNodes     []Node
)

func CreateBlockChain() error {
	return nil
}

func CreateObserver() error {
	return nil
}

func ReleaseObserver(pubKey types.PubKey) error {
	return nil
}

func ReleaseAllObservers() error {
	return nil
}

func DestroyBlockChain() error {
	return nil
}

func GetGenesisNodes() []Node {
	return nil
}
