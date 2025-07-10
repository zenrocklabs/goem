package goem

import (
	"math/big"
	"testing"

	"github.com/VenimirPetkov/goem/ethereum"
	"github.com/VenimirPetkov/goem/solana"
)

func TestEthereumNetworks(t *testing.T) {
	// Test that all Ethereum networks have valid chain IDs
	networks := []struct {
		name     string
		chainId  *big.Int
		expected *big.Int
	}{
		{"Mainnet", ethereum.MainnetChainId, big.NewInt(1)},
		{"Holesky", ethereum.HoleskyChainId, big.NewInt(17000)},
		{"Hoodi", ethereum.HoodiChainId, big.NewInt(560048)},
	}

	for _, network := range networks {
		if network.chainId.Cmp(network.expected) != 0 {
			t.Errorf("%s chain ID = %v, want %v", network.name, network.chainId, network.expected)
		}
	}
}

func TestSolanaNetworks(t *testing.T) {
	// Test that all Solana networks have valid genesis hashes
	networks := []struct {
		name         string
		genesisHash  string
		expectedHash string
	}{
		{"Mainnet", solana.MainnetGenesisHash, "5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d"},
		{"Devnet", solana.DevnetGenesisHash, "EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG"},
		{"Testnet", solana.TestnetGenesisHash, "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY"},
	}

	for _, network := range networks {
		if network.genesisHash != network.expectedHash {
			t.Errorf("%s genesis hash = %s, want %s", network.name, network.genesisHash, network.expectedHash)
		}
	}
}

func TestCAIP2Identifiers(t *testing.T) {
	// Test Ethereum CAIP-2 identifiers
	ethereumNetworks := []struct {
		name     string
		caip2    string
		expected string
	}{
		{"Mainnet", ethereum.MainnetCAIP2, "eip155:1"},
		{"Holesky", ethereum.HoleskyCAIP2, "eip155:17000"},
		{"Hoodi", ethereum.HoodiCAIP2, "eip155:560048"},
	}

	for _, network := range ethereumNetworks {
		if network.caip2 != network.expected {
			t.Errorf("%s CAIP-2 = %s, want %s", network.name, network.caip2, network.expected)
		}
	}

	// Test Solana CAIP-2 identifiers
	solanaNetworks := []struct {
		name     string
		caip2    string
		expected string
	}{
		{"Mainnet", solana.MainnetCAIP2, "solana:5eykt4UsFv8P8NJdTREpY1vzqKqZKvdp"},
		{"Devnet", solana.DevnetCAIP2, "solana:EtWTRABZaYq6iMfeYKouRu166VU2xqa1"},
		{"Testnet", solana.TestnetCAIP2, "solana:4fYNw3dojWmQnv4KTNLWqTfotHJgf6YQePvCYTb3RqaU"},
	}

	for _, network := range solanaNetworks {
		if network.caip2 != network.expected {
			t.Errorf("%s CAIP-2 = %s, want %s", network.name, network.caip2, network.expected)
		}
	}
}

func TestAllURLsAreValid(t *testing.T) {
	// Test Ethereum URLs
	ethereumURLs := []struct {
		name string
		url  string
	}{
		{"Mainnet RPC", ethereum.MainnetRPC},
		{"Mainnet WS", ethereum.MainnetWS},
		{"Mainnet Beacon", ethereum.MainnetBeacon},
		{"Mainnet Explorer", ethereum.MainnetExplorer},
		{"Mainnet API", ethereum.MainnetAPI},
		{"Holesky RPC", ethereum.HoleskyRPC},
		{"Holesky WS", ethereum.HoleskyWS},
		{"Holesky Beacon", ethereum.HoleskyBeacon},
		{"Holesky Explorer", ethereum.HoleskyExplorer},
		{"Holesky API", ethereum.HoleskyAPI},
		{"Hoodi RPC", ethereum.HoodiRPC},
		{"Hoodi WS", ethereum.HoodiWS},
		{"Hoodi Beacon", ethereum.HoodiBeacon},
		{"Hoodi Explorer", ethereum.HoodiExplorer},
		{"Hoodi API", ethereum.HoodiAPI},
	}

	for _, url := range ethereumURLs {
		if url.url == "" {
			t.Errorf("%s is empty", url.name)
		}
	}

	// Test Solana URLs
	solanaURLs := []struct {
		name string
		url  string
	}{
		{"Mainnet RPC", solana.MainnetRPC},
		{"Mainnet WS", solana.MainnetWS},
		{"Mainnet Explorer", solana.MainnetExplorer},
		{"Mainnet API", solana.MainnetAPI},
		{"Devnet RPC", solana.DevnetRPC},
		{"Devnet WS", solana.DevnetWS},
		{"Devnet Explorer", solana.DevnetExplorer},
		{"Devnet API", solana.DevnetAPI},
		{"Testnet RPC", solana.TestnetRPC},
		{"Testnet WS", solana.TestnetWS},
		{"Testnet Explorer", solana.TestnetExplorer},
		{"Testnet API", solana.TestnetAPI},
	}

	for _, url := range solanaURLs {
		if url.url == "" {
			t.Errorf("%s is empty", url.name)
		}
	}
}

func TestPackageStructure(t *testing.T) {
	// Test that all expected variables are exported and accessible
	_ = ethereum.MainnetChainId
	_ = ethereum.MainnetCAIP2
	_ = ethereum.MainnetRPC
	_ = ethereum.MainnetWS
	_ = ethereum.MainnetBeacon
	_ = ethereum.MainnetExplorer
	_ = ethereum.MainnetAPI

	_ = ethereum.HoleskyChainId
	_ = ethereum.HoleskyCAIP2
	_ = ethereum.HoleskyRPC
	_ = ethereum.HoleskyWS
	_ = ethereum.HoleskyBeacon
	_ = ethereum.HoleskyExplorer
	_ = ethereum.HoleskyAPI

	_ = ethereum.HoodiChainId
	_ = ethereum.HoodiCAIP2
	_ = ethereum.HoodiRPC
	_ = ethereum.HoodiWS
	_ = ethereum.HoodiBeacon
	_ = ethereum.HoodiExplorer
	_ = ethereum.HoodiAPI

	_ = solana.MainnetGenesisHash
	_ = solana.MainnetCAIP2
	_ = solana.MainnetRPC
	_ = solana.MainnetWS
	_ = solana.MainnetExplorer
	_ = solana.MainnetAPI

	_ = solana.DevnetGenesisHash
	_ = solana.DevnetCAIP2
	_ = solana.DevnetRPC
	_ = solana.DevnetWS
	_ = solana.DevnetExplorer
	_ = solana.DevnetAPI

	_ = solana.TestnetGenesisHash
	_ = solana.TestnetCAIP2
	_ = solana.TestnetRPC
	_ = solana.TestnetWS
	_ = solana.TestnetExplorer
	_ = solana.TestnetAPI
}
