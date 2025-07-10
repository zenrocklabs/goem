package solana

import "testing"

func TestMainnetConfiguration(t *testing.T) {
	// Test genesis hash
	expectedGenesisHash := "5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d"
	if MainnetGenesisHash != expectedGenesisHash {
		t.Errorf("MainnetGenesisHash = %s, want %s", MainnetGenesisHash, expectedGenesisHash)
	}

	// Test CAIP-2 identifier
	expectedCAIP2 := "solana:5eykt4UsFv8P8NJdTREpY1vzqKqZKvdp"
	if MainnetCAIP2 != expectedCAIP2 {
		t.Errorf("MainnetCAIP2 = %s, want %s", MainnetCAIP2, expectedCAIP2)
	}

	// Test RPC URL
	expectedRPC := "https://api.mainnet-beta.solana.com"
	if MainnetRPC != expectedRPC {
		t.Errorf("MainnetRPC = %s, want %s", MainnetRPC, expectedRPC)
	}

	// Test WebSocket URL
	expectedWS := "wss://api.mainnet-beta.solana.com"
	if MainnetWS != expectedWS {
		t.Errorf("MainnetWS = %s, want %s", MainnetWS, expectedWS)
	}

	// Test Explorer URL
	expectedExplorer := "https://explorer.solana.com"
	if MainnetExplorer != expectedExplorer {
		t.Errorf("MainnetExplorer = %s, want %s", MainnetExplorer, expectedExplorer)
	}

	// Test API URL
	expectedAPI := "https://api.solscan.io"
	if MainnetAPI != expectedAPI {
		t.Errorf("MainnetAPI = %s, want %s", MainnetAPI, expectedAPI)
	}

	// Test Network Name
	expectedNetworkName := "mainnet"
	if MainnetNetworkName != expectedNetworkName {
		t.Errorf("MainnetNetworkName = %s, want %s", MainnetNetworkName, expectedNetworkName)
	}

	// Test Environment
	expectedEnvironment := "prod"
	if MainnetEnvironment != expectedEnvironment {
		t.Errorf("MainnetEnvironment = %s, want %s", MainnetEnvironment, expectedEnvironment)
	}
}

func TestMainnetURLsAreValid(t *testing.T) {
	// Test that URLs are not empty
	if MainnetRPC == "" {
		t.Error("MainnetRPC is empty")
	}
	if MainnetWS == "" {
		t.Error("MainnetWS is empty")
	}
	if MainnetExplorer == "" {
		t.Error("MainnetExplorer is empty")
	}
	if MainnetAPI == "" {
		t.Error("MainnetAPI is empty")
	}

	// Test that URLs start with expected protocols
	if MainnetRPC[:8] != "https://" {
		t.Errorf("MainnetRPC should start with https://, got %s", MainnetRPC[:8])
	}
	if MainnetWS[:6] != "wss://" {
		t.Errorf("MainnetWS should start with wss://, got %s", MainnetWS[:6])
	}
	if MainnetExplorer[:8] != "https://" {
		t.Errorf("MainnetExplorer should start with https://, got %s", MainnetExplorer[:8])
	}
	if MainnetAPI[:8] != "https://" {
		t.Errorf("MainnetAPI should start with https://, got %s", MainnetAPI[:8])
	}
}
