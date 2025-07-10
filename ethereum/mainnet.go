package ethereum

import "math/big"

// Mainnet configuration
var (
	MainnetChainId  *big.Int = big.NewInt(1)
	MainnetCAIP2    string   = "eip155:1"
	MainnetRPC      string   = "https://ethereum-rpc.publicnode.com"
	MainnetWS       string   = "wss://ethereum-rpc.publicnode.com"
	MainnetBeacon   string   = "https://ethereum-beacon-api.publicnode.com"
	MainnetExplorer string   = "https://etherscan.io"
	MainnetAPI      string   = "https://api.etherscan.io"
)
