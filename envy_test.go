package envy

import (
	"os"
	"testing"
)

func TestGetEnvNoEnv(t *testing.T) {
	result := GetEnv("TEST_KEY", "FALLBACK")
	if result != "FALLBACK" {
		t.Error("Expected FALLBACK received", result)
	}
}

func TestGetEnvWithEnv(t *testing.T) {
	os.Setenv("TEST_KEY_X", "ENV")
	if os.Getenv("TEST_KEY_X") != "ENV" {
		t.Error("Could not properly set environment variable TEST_KEY_X")
	}
	result := GetEnv("TEST_KEY_X", "FALLBACK")
	if result != "ENV" {
		t.Error("Expected ENV received", result)
	}
}
