package ethereum

import (
	"math/big"
	"testing"
)

func TestHoodiConfiguration(t *testing.T) {
	// Test chain ID
	expectedChainId := big.NewInt(560048)
	if HoodiChainId.Cmp(expectedChainId) != 0 {
		t.Errorf("HoodiChainId = %v, want %v", HoodiChainId, expectedChainId)
	}

	// Test CAIP-2 identifier
	expectedCAIP2 := "eip155:560048"
	if HoodiCAIP2 != expectedCAIP2 {
		t.Errorf("HoodiCAIP2 = %s, want %s", HoodiCAIP2, expectedCAIP2)
	}

	// Test RPC URL
	expectedRPC := "https://ethereum-hoodi-rpc.publicnode.com"
	if HoodiRPC != expectedRPC {
		t.Errorf("HoodiRPC = %s, want %s", HoodiRPC, expectedRPC)
	}

	// Test WebSocket URL
	expectedWS := "wss://ethereum-hoodi-rpc.publicnode.com"
	if HoodiWS != expectedWS {
		t.Errorf("HoodiWS = %s, want %s", HoodiWS, expectedWS)
	}

	// Test Beacon API URL
	expectedBeacon := "https://ethereum-beacon-api.publicnode.com"
	if HoodiBeacon != expectedBeacon {
		t.Errorf("HoodiBeacon = %s, want %s", HoodiBeacon, expectedBeacon)
	}

	// Test Explorer URL
	expectedExplorer := "https://hoodi.etherscan.io"
	if HoodiExplorer != expectedExplorer {
		t.Errorf("HoodiExplorer = %s, want %s", HoodiExplorer, expectedExplorer)
	}

	// Test API URL
	expectedAPI := "https://api-hoodi.etherscan.io"
	if HoodiAPI != expectedAPI {
		t.Errorf("HoodiAPI = %s, want %s", HoodiAPI, expectedAPI)
	}
}

func TestHoodiURLsAreValid(t *testing.T) {
	// Test that URLs are not empty
	if HoodiRPC == "" {
		t.Error("HoodiRPC is empty")
	}
	if HoodiWS == "" {
		t.Error("HoodiWS is empty")
	}
	if HoodiBeacon == "" {
		t.Error("HoodiBeacon is empty")
	}
	if HoodiExplorer == "" {
		t.Error("HoodiExplorer is empty")
	}
	if HoodiAPI == "" {
		t.Error("HoodiAPI is empty")
	}

	// Test that URLs start with expected protocols
	if HoodiRPC[:8] != "https://" {
		t.Errorf("HoodiRPC should start with https://, got %s", HoodiRPC[:8])
	}
	if HoodiWS[:6] != "wss://" {
		t.Errorf("HoodiWS should start with wss://, got %s", HoodiWS[:6])
	}
	if HoodiBeacon[:8] != "https://" {
		t.Errorf("HoodiBeacon should start with https://, got %s", HoodiBeacon[:8])
	}
	if HoodiExplorer[:8] != "https://" {
		t.Errorf("HoodiExplorer should start with https://, got %s", HoodiExplorer[:8])
	}
	if HoodiAPI[:8] != "https://" {
		t.Errorf("HoodiAPI should start with https://, got %s", HoodiAPI[:8])
	}
}
