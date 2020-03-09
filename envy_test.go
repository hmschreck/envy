package envy

import (
	"os"
	"testing"
	"github.com/magiconair/properties/assert"
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

func TestGetEnvInt(t *testing.T) {
	os.Setenv("TEST_INT", "24")
	result, err := GetEnvInt("TEST_INT", "23")
	if err != nil {
		t.Error("GOt error when shouldn't have")
	}
	assert.Equal(t, result, 24)
}

func TestGetEnvIntFallback(t *testing.T) {
	result, err := GetEnvInt("TEST_INT_2", "23")
	if err != nil {
		t.Error("Got error when shouldn't have")
	}
	assert.Equal(t, result, 23)
}

func TestGetEnvIntNonInt(t *testing.T) {
	_, err := GetEnvInt("TEST_INT_NONE", "asdf")
	if err == nil {
		t.Error("Did not receive error when should have")
	}
}
