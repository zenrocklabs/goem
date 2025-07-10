package ethereum

import "math/big"

// Hoodi testnet configuration
var (
	HoodiChainId     *big.Int = big.NewInt(560048)
	HoodiCAIP2       string   = "eip155:560048"
	HoodiRPC         string   = "https://ethereum-hoodi-rpc.publicnode.com"
	HoodiWS          string   = "wss://ethereum-hoodi-rpc.publicnode.com"
	HoodiBeacon      string   = "https://ethereum-beacon-api.publicnode.com"
	HoodiExplorer    string   = "https://hoodi.etherscan.io"
	HoodiAPI         string   = "https://api-hoodi.etherscan.io"
	HoodiNetworkName string   = "hoodi"
	HoodiEnvironment string   = "testnet"
)
