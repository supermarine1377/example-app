package config_test

import (
	"testing"

	"github.com/supermarine1377/example-app/config"
)

func TestNew(t *testing.T) {
	wantPort := "80"
	cfg, err := config.New()
	if err != nil {
		t.Errorf("error: failed to get config: %+v", err)
	}
	if cfg.Port != wantPort {
		t.Errorf("error: unexpected port: expected %s, but got %s", wantPort, cfg.Port)
	}
	wantEnv := "dev"
	if cfg.Env != wantEnv {
		t.Errorf("error: unexpected dev: expected %s, but got %s", wantEnv, cfg.Env)
	}
}
