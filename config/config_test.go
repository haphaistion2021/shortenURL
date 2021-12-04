package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	config := &Config
	if config == nil {
		t.Fatalf("Expected no nil, got nil")
	}
}
