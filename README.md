# GoEM - Ethereum and Solana Configuration

This repository contains centralized configuration for Ethereum and Solana networks including chain IDs, CAIP-2 identifiers, public RPC URLs, WebSocket URLs, beacon API URLs, block explorer URLs, and API endpoints.

## Structure

```
goem/
├── ethereum/
│   ├── constants.go   # Common constants (CAIP-2 namespace)
│   ├── mainnet.go     # Ethereum mainnet (chain ID: 1)
│   ├── holesky.go     # Ethereum Holesky testnet (chain ID: 17000)
│   ├── hoodi.go       # Ethereum Hoodi testnet (chain ID: 560048)
│   └── constants_test.go
└── solana/
    ├── constants.go    # Common constants (CAIP-2 namespace)
    ├── mainnet.go      # Solana mainnet (genesis hash)
    ├── devnet.go       # Solana devnet (genesis hash)
    ├── testnet.go      # Solana testnet (genesis hash)
    └── constants_test.go
```

## Usage

Import the package and use the constants:

```go
import (
    "github.com/VenimirPetkov/goem/ethereum"
    "github.com/VenimirPetkov/goem/solana"
)

// Common namespace constants
ethereumNamespace := ethereum.CAIP2Namespace  // "eip155"
solanaNamespace := solana.CAIP2Namespace      // "solana"

// Ethereum (chain IDs are *big.Int to match ethclient.ChainID())
chainId := ethereum.HoodiChainId
caip2 := ethereum.HoodiCAIP2
rpc := ethereum.HoodiRPC
ws := ethereum.HoodiWS
beacon := ethereum.HoodiBeacon
explorer := ethereum.HoodiExplorer
api := ethereum.HoodiAPI
networkName := ethereum.HoodiNetworkName
environment := ethereum.HoodiEnvironment

// Solana (uses genesis hashes)
genesisHash := solana.MainnetGenesisHash
caip2 := solana.MainnetCAIP2
rpc := solana.MainnetRPC
ws := solana.MainnetWS
explorer := solana.MainnetExplorer
api := solana.MainnetAPI
networkName := solana.MainnetNetworkName
environment := solana.MainnetEnvironment
```

## Networks

### Ethereum
- **Mainnet**: Chain ID 1, CAIP-2 `eip155:1`
  - RPC: `https://ethereum-rpc.publicnode.com`
  - WS: `wss://ethereum-rpc.publicnode.com`
  - Beacon: `https://ethereum-beacon-api.publicnode.com`
  - Explorer: `https://etherscan.io`
  - API: `https://api.etherscan.io`
  - Network Name: `mainnet`
  - Environment: `prod`
- **Holesky**: Chain ID 17000, CAIP-2 `eip155:17000`
  - RPC: `https://ethereum-holesky-rpc.publicnode.com`
  - WS: `wss://ethereum-holesky-rpc.publicnode.com`
  - Beacon: `https://ethereum-beacon-api.publicnode.com`
  - Explorer: `https://holesky.etherscan.io`
  - API: `https://api-holesky.etherscan.io`
  - Network Name: `holesky`
  - Environment: `testnet`
- **Hoodi**: Chain ID 560048, CAIP-2 `eip155:560048`
  - RPC: `https://ethereum-hoodi-rpc.publicnode.com`
  - WS: `wss://ethereum-hoodi-rpc.publicnode.com`
  - Beacon: `https://ethereum-beacon-api.publicnode.com`
  - Explorer: `https://hoodi.etherscan.io`
  - API: `https://api-hoodi.etherscan.io`
  - Network Name: `hoodi`
  - Environment: `testnet`

### Solana
- **Mainnet**: Genesis Hash `5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d`, CAIP-2 `solana:5eykt4UsFv8P8NJdTREpY1vzqKqZKvdp`
  - RPC: `https://api.mainnet-beta.solana.com`
  - WS: `wss://api.mainnet-beta.solana.com`
  - Explorer: `https://explorer.solana.com`
  - API: `https://api.solscan.io`
  - Network Name: `mainnet`
  - Environment: `prod`
- **Devnet**: Genesis Hash `EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG`, CAIP-2 `solana:EtWTRABZaYq6iMfeYKouRu166VU2xqa1`
  - RPC: `https://api.devnet.solana.com`
  - WS: `wss://api.devnet.solana.com`
  - Explorer: `https://explorer.solana.com/?cluster=devnet`
  - API: `https://api.solscan.io?cluster=devnet`
  - Network Name: `devnet`
  - Environment: `devnet`
- **Testnet**: Genesis Hash `4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY`, CAIP-2 `solana:4fYNw3dojWmQnv4KTNLWqTfotHJgf6YQePvCYTb3RqaU`
  - RPC: `https://api.testnet.solana.com`
  - WS: `wss://api.testnet.solana.com`
  - Explorer: `https://explorer.solana.com/?cluster=testnet`
  - API: `https://api.solscan.io?cluster=testnet`
  - Network Name: `testnet`
  - Environment: `testnet`

## Important Notes

### CAIP-2 Namespaces
The package provides namespace constants for CAIP-2 format:
- `ethereum.CAIP2Namespace` = `"eip155"` for Ethereum chains
- `solana.CAIP2Namespace` = `"solana"` for Solana chains

These constants can be used when validating CAIP-2 strings or building CAIP-2 identifiers programmatically.

### Ethereum Chain IDs
Ethereum chain IDs are represented as `*big.Int` to match the return type of `ethclient.ChainID()`. This allows for direct comparison with chain IDs returned by the Go Ethereum client.

### Solana Genesis Hashes
Solana does not use EVM-style integer chain IDs. Instead, each cluster is distinguished by its genesis hash, which is the cryptographic fingerprint of the very first block's data. To verify the connection to the correct cluster, a client can fetch `getGenesisHash` from the RPC node and compare it to the known hash.

### Solana Transactions
Solana transactions do not carry a numeric chainId. When building and signing a transaction, a recent blockhash is included instead. Solana relies on blockhash uniqueness per cluster to prevent cross-network replay.

## Contribution

Anyone is welcome to add new networks! Simply follow the existing pattern for Ethereum or Solana in the appropriate directory, add a test file for your network, and open a Pull Request (PR).

- For Ethereum: add a new file in `goem/ethereum/` following the structure of `mainnet.go`, `holesky.go`, or `hoodi.go`.
- For Solana: add a new file in `goem/solana/` following the structure of `mainnet.go`, `devnet.go`, or `testnet.go`.
- Please include a corresponding test file to verify your configuration.
- Update the README if needed.

We welcome your contributions!