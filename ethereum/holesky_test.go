package ethereum

import (
	"math/big"
	"testing"
)

func TestHoleskyConfiguration(t *testing.T) {
	// Test chain ID
	expectedChainId := big.NewInt(17000)
	if HoleskyChainId.Cmp(expectedChainId) != 0 {
		t.Errorf("HoleskyChainId = %v, want %v", HoleskyChainId, expectedChainId)
	}

	// Test CAIP-2 identifier
	expectedCAIP2 := "eip155:17000"
	if HoleskyCAIP2 != expectedCAIP2 {
		t.Errorf("HoleskyCAIP2 = %s, want %s", HoleskyCAIP2, expectedCAIP2)
	}

	// Test RPC URL
	expectedRPC := "https://ethereum-holesky-rpc.publicnode.com"
	if HoleskyRPC != expectedRPC {
		t.Errorf("HoleskyRPC = %s, want %s", HoleskyRPC, expectedRPC)
	}

	// Test WebSocket URL
	expectedWS := "wss://ethereum-holesky-rpc.publicnode.com"
	if HoleskyWS != expectedWS {
		t.Errorf("HoleskyWS = %s, want %s", HoleskyWS, expectedWS)
	}

	// Test Beacon API URL
	expectedBeacon := "https://ethereum-beacon-api.publicnode.com"
	if HoleskyBeacon != expectedBeacon {
		t.Errorf("HoleskyBeacon = %s, want %s", HoleskyBeacon, expectedBeacon)
	}

	// Test Explorer URL
	expectedExplorer := "https://holesky.etherscan.io"
	if HoleskyExplorer != expectedExplorer {
		t.Errorf("HoleskyExplorer = %s, want %s", HoleskyExplorer, expectedExplorer)
	}

	// Test API URL
	expectedAPI := "https://api-holesky.etherscan.io"
	if HoleskyAPI != expectedAPI {
		t.Errorf("HoleskyAPI = %s, want %s", HoleskyAPI, expectedAPI)
	}

	// Test Network Name
	expectedNetworkName := "holesky"
	if HoleskyNetworkName != expectedNetworkName {
		t.Errorf("HoleskyNetworkName = %s, want %s", HoleskyNetworkName, expectedNetworkName)
	}

	// Test Environment
	expectedEnvironment := "testnet"
	if HoleskyEnvironment != expectedEnvironment {
		t.Errorf("HoleskyEnvironment = %s, want %s", HoleskyEnvironment, expectedEnvironment)
	}
}

func TestHoleskyURLsAreValid(t *testing.T) {
	// Test that URLs are not empty
	if HoleskyRPC == "" {
		t.Error("HoleskyRPC is empty")
	}
	if HoleskyWS == "" {
		t.Error("HoleskyWS is empty")
	}
	if HoleskyBeacon == "" {
		t.Error("HoleskyBeacon is empty")
	}
	if HoleskyExplorer == "" {
		t.Error("HoleskyExplorer is empty")
	}
	if HoleskyAPI == "" {
		t.Error("HoleskyAPI is empty")
	}

	// Test that URLs start with expected protocols
	if HoleskyRPC[:8] != "https://" {
		t.Errorf("HoleskyRPC should start with https://, got %s", HoleskyRPC[:8])
	}
	if HoleskyWS[:6] != "wss://" {
		t.Errorf("HoleskyWS should start with wss://, got %s", HoleskyWS[:6])
	}
	if HoleskyBeacon[:8] != "https://" {
		t.Errorf("HoleskyBeacon should start with https://, got %s", HoleskyBeacon[:8])
	}
	if HoleskyExplorer[:8] != "https://" {
		t.Errorf("HoleskyExplorer should start with https://, got %s", HoleskyExplorer[:8])
	}
	if HoleskyAPI[:8] != "https://" {
		t.Errorf("HoleskyAPI should start with https://, got %s", HoleskyAPI[:8])
	}
}
