package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
)

// GetRPCApis returns a list of RPC APIs for testing
func GetRPCApis(t *testing.T, apiBackend ethapi.Backend) []rpc.API {
	apis := ethapi.GetAPIs(apiBackend)
	apis = append(apis, tracers.APIs(nil)...)
	return apis
}
