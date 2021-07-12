package config

import (
    "testing"
)

func TestFlags(t *testing.T) {
	expected_result := "vmware_exporter.yml"
    f := ParseFlags()
	if f.ConfigFile != expected_result {
		t.Fatalf("Expected configfile to contain \"%s\" but got \"%s\".", expected_result, f.ConfigFile)
	}
}

func TestConfig(t *testing.T) {
	expected_api_url := "https://localhost"
	expected_api_userid := "userid"
	expected_api_password := "password"
	expected_api_insecure := true
	cfg, err := ParseConfig("vmware_exporter.yml")
	if err != nil {
		t.Fatalf("newConfig returned: %v", err)
	}
	if cfg.API.URL != expected_api_url {
		t.Fatalf("Expected cfg.API.URL to contain \"%s\" but got \"%s\".", expected_api_url, cfg.API.URL)
	}
	if cfg.API.UserID != expected_api_userid {
		t.Fatalf("Expected cfg.API.UserID to contain \"%s\" but got \"%s\".", expected_api_userid, cfg.API.UserID)
	}
	if cfg.API.Password != expected_api_password {
		t.Fatalf("Expected cfg.API.Password to contain \"%s\" but got \"%s\".", expected_api_password, cfg.API.Password)
	}
	if cfg.API.Insecure != expected_api_insecure {
		t.Fatalf("Expected cfg.API.Insecure to be \"%t\" but got \"%t\".", expected_api_insecure, cfg.API.Insecure)
	}
}