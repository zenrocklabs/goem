package ethereum

import "math/big"

// Holesky testnet configuration
var (
	HoleskyChainId  *big.Int = big.NewInt(17000)
	HoleskyCAIP2    string   = "eip155:17000"
	HoleskyRPC      string   = "https://ethereum-holesky-rpc.publicnode.com"
	HoleskyWS       string   = "wss://ethereum-holesky-rpc.publicnode.com"
	HoleskyBeacon   string   = "https://ethereum-beacon-api.publicnode.com"
	HoleskyExplorer string   = "https://holesky.etherscan.io"
	HoleskyAPI      string   = "https://api-holesky.etherscan.io"
)
