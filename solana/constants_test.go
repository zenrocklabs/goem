package solana

import (
	"testing"
)

func TestCAIP2Namespace(t *testing.T) {
	expected := "solana"
	if CAIP2Namespace != expected {
		t.Errorf("Expected CAIP2Namespace to be %s, got %s", expected, CAIP2Namespace)
	}
}
