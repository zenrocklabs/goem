package ethereum

import (
	"math/big"
	"testing"
)

func TestMainnetConfiguration(t *testing.T) {
	// Test chain ID
	expectedChainId := big.NewInt(1)
	if MainnetChainId.Cmp(expectedChainId) != 0 {
		t.Errorf("MainnetChainId = %v, want %v", MainnetChainId, expectedChainId)
	}

	// Test CAIP-2 identifier
	expectedCAIP2 := "eip155:1"
	if MainnetCAIP2 != expectedCAIP2 {
		t.Errorf("MainnetCAIP2 = %s, want %s", MainnetCAIP2, expectedCAIP2)
	}

	// Test RPC URL
	expectedRPC := "https://ethereum-rpc.publicnode.com"
	if MainnetRPC != expectedRPC {
		t.Errorf("MainnetRPC = %s, want %s", MainnetRPC, expectedRPC)
	}

	// Test WebSocket URL
	expectedWS := "wss://ethereum-rpc.publicnode.com"
	if MainnetWS != expectedWS {
		t.Errorf("MainnetWS = %s, want %s", MainnetWS, expectedWS)
	}

	// Test Beacon API URL
	expectedBeacon := "https://ethereum-beacon-api.publicnode.com"
	if MainnetBeacon != expectedBeacon {
		t.Errorf("MainnetBeacon = %s, want %s", MainnetBeacon, expectedBeacon)
	}

	// Test Explorer URL
	expectedExplorer := "https://etherscan.io"
	if MainnetExplorer != expectedExplorer {
		t.Errorf("MainnetExplorer = %s, want %s", MainnetExplorer, expectedExplorer)
	}

	// Test API URL
	expectedAPI := "https://api.etherscan.io"
	if MainnetAPI != expectedAPI {
		t.Errorf("MainnetAPI = %s, want %s", MainnetAPI, expectedAPI)
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
	if MainnetBeacon == "" {
		t.Error("MainnetBeacon is empty")
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
	if MainnetBeacon[:8] != "https://" {
		t.Errorf("MainnetBeacon should start with https://, got %s", MainnetBeacon[:8])
	}
	if MainnetExplorer[:8] != "https://" {
		t.Errorf("MainnetExplorer should start with https://, got %s", MainnetExplorer[:8])
	}
	if MainnetAPI[:8] != "https://" {
		t.Errorf("MainnetAPI should start with https://, got %s", MainnetAPI[:8])
	}
}
