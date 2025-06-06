package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	jsonContent := `{
        "port": 8080,
        "routes": [
            {
                "path_prefix": "/api",
                "target": "http://localhost:3000"
            },
            {	
                "path_prefix": "/static",
                "target": "http://localhost:8081"
            }
        ]
    }`

	tmpFile, err := os.CreateTemp("", "config_test_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	tmpFile.WriteString(jsonContent)
	tmpFile.Close()

	cfg, err := Load(tmpFile.Name())
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	if cfg.Port != 8080 {
		t.Errorf("Expected Port to be 8080, got %q", cfg.Port)
	}

	if len(cfg.Routes) != 2 {
		t.Fatalf("Expected 2 routes, got %d", len(cfg.Routes))
	}

	if cfg.Routes[0].PathPrefix != "/api" || cfg.Routes[0].Target != "http://localhost:3000" {
		t.Errorf("First route mismatch: %+v", cfg.Routes[0])
	}

	if cfg.Routes[1].PathPrefix != "/static" || cfg.Routes[1].Target != "http://localhost:8081" {
		t.Errorf("Second route mismatch: %+v", cfg.Routes[1])
	}
}
