package solana

import "testing"

func TestTestnetConfiguration(t *testing.T) {
	// Test genesis hash
	expectedGenesisHash := "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY"
	if TestnetGenesisHash != expectedGenesisHash {
		t.Errorf("TestnetGenesisHash = %s, want %s", TestnetGenesisHash, expectedGenesisHash)
	}

	// Test CAIP-2 identifier
	expectedCAIP2 := "solana:4fYNw3dojWmQnv4KTNLWqTfotHJgf6YQePvCYTb3RqaU"
	if TestnetCAIP2 != expectedCAIP2 {
		t.Errorf("TestnetCAIP2 = %s, want %s", TestnetCAIP2, expectedCAIP2)
	}

	// Test RPC URL
	expectedRPC := "https://api.testnet.solana.com"
	if TestnetRPC != expectedRPC {
		t.Errorf("TestnetRPC = %s, want %s", TestnetRPC, expectedRPC)
	}

	// Test WebSocket URL
	expectedWS := "wss://api.testnet.solana.com"
	if TestnetWS != expectedWS {
		t.Errorf("TestnetWS = %s, want %s", TestnetWS, expectedWS)
	}

	// Test Explorer URL
	expectedExplorer := "https://explorer.solana.com/?cluster=testnet"
	if TestnetExplorer != expectedExplorer {
		t.Errorf("TestnetExplorer = %s, want %s", TestnetExplorer, expectedExplorer)
	}

	// Test API URL
	expectedAPI := "https://api.solscan.io?cluster=testnet"
	if TestnetAPI != expectedAPI {
		t.Errorf("TestnetAPI = %s, want %s", TestnetAPI, expectedAPI)
	}
}

func TestTestnetURLsAreValid(t *testing.T) {
	// Test that URLs are not empty
	if TestnetRPC == "" {
		t.Error("TestnetRPC is empty")
	}
	if TestnetWS == "" {
		t.Error("TestnetWS is empty")
	}
	if TestnetExplorer == "" {
		t.Error("TestnetExplorer is empty")
	}
	if TestnetAPI == "" {
		t.Error("TestnetAPI is empty")
	}

	// Test that URLs start with expected protocols
	if TestnetRPC[:8] != "https://" {
		t.Errorf("TestnetRPC should start with https://, got %s", TestnetRPC[:8])
	}
	if TestnetWS[:6] != "wss://" {
		t.Errorf("TestnetWS should start with wss://, got %s", TestnetWS[:6])
	}
	if TestnetExplorer[:8] != "https://" {
		t.Errorf("TestnetExplorer should start with https://, got %s", TestnetExplorer[:8])
	}
	if TestnetAPI[:8] != "https://" {
		t.Errorf("TestnetAPI should start with https://, got %s", TestnetAPI[:8])
	}
}
