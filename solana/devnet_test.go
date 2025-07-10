package solana

import "testing"

func TestDevnetConfiguration(t *testing.T) {
	// Test genesis hash
	expectedGenesisHash := "EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG"
	if DevnetGenesisHash != expectedGenesisHash {
		t.Errorf("DevnetGenesisHash = %s, want %s", DevnetGenesisHash, expectedGenesisHash)
	}

	// Test CAIP-2 identifier
	expectedCAIP2 := "solana:EtWTRABZaYq6iMfeYKouRu166VU2xqa1"
	if DevnetCAIP2 != expectedCAIP2 {
		t.Errorf("DevnetCAIP2 = %s, want %s", DevnetCAIP2, expectedCAIP2)
	}

	// Test RPC URL
	expectedRPC := "https://api.devnet.solana.com"
	if DevnetRPC != expectedRPC {
		t.Errorf("DevnetRPC = %s, want %s", DevnetRPC, expectedRPC)
	}

	// Test WebSocket URL
	expectedWS := "wss://api.devnet.solana.com"
	if DevnetWS != expectedWS {
		t.Errorf("DevnetWS = %s, want %s", DevnetWS, expectedWS)
	}

	// Test Explorer URL
	expectedExplorer := "https://explorer.solana.com/?cluster=devnet"
	if DevnetExplorer != expectedExplorer {
		t.Errorf("DevnetExplorer = %s, want %s", DevnetExplorer, expectedExplorer)
	}

	// Test API URL
	expectedAPI := "https://api.solscan.io?cluster=devnet"
	if DevnetAPI != expectedAPI {
		t.Errorf("DevnetAPI = %s, want %s", DevnetAPI, expectedAPI)
	}

	// Test Network Name
	expectedNetworkName := "devnet"
	if DevnetNetworkName != expectedNetworkName {
		t.Errorf("DevnetNetworkName = %s, want %s", DevnetNetworkName, expectedNetworkName)
	}

	// Test Environment
	expectedEnvironment := "devnet"
	if DevnetEnvironment != expectedEnvironment {
		t.Errorf("DevnetEnvironment = %s, want %s", DevnetEnvironment, expectedEnvironment)
	}
}

func TestDevnetURLsAreValid(t *testing.T) {
	// Test that URLs are not empty
	if DevnetRPC == "" {
		t.Error("DevnetRPC is empty")
	}
	if DevnetWS == "" {
		t.Error("DevnetWS is empty")
	}
	if DevnetExplorer == "" {
		t.Error("DevnetExplorer is empty")
	}
	if DevnetAPI == "" {
		t.Error("DevnetAPI is empty")
	}

	// Test that URLs start with expected protocols
	if DevnetRPC[:8] != "https://" {
		t.Errorf("DevnetRPC should start with https://, got %s", DevnetRPC[:8])
	}
	if DevnetWS[:6] != "wss://" {
		t.Errorf("DevnetWS should start with wss://, got %s", DevnetWS[:6])
	}
	if DevnetExplorer[:8] != "https://" {
		t.Errorf("DevnetExplorer should start with https://, got %s", DevnetExplorer[:8])
	}
	if DevnetAPI[:8] != "https://" {
		t.Errorf("DevnetAPI should start with https://, got %s", DevnetAPI[:8])
	}
}
