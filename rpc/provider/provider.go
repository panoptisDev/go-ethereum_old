package provider

import "github.com/ethereum/go-ethereum/internal/ethapi"

func GetLegacyEthAPI(b ethapi.Backend) *ethapi.BlockChainAPI {
	return ethapi.NewBlockChainAPI(b)
}
