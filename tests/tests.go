package tests

import (
	"api/internal/config"
	"testing"
)

func TestEnv(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Errorf("cfg make err %s", cfg)
	}

}
