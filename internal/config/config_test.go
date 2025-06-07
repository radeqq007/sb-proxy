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
        ],
				"headers": {
					"add": {
						"X-Proxy-By": "sb-proxy"
					},
					"remove": ["X-Powered-By"]
				},
				"timeout_ms": 5000
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

	if len(cfg.Headers.Add) != 1 || cfg.Headers.Add["X-Proxy-By"] != "sb-proxy" {
		t.Errorf("Expected header 'X-Proxy-By' to be 'sb-proxy', got %q", cfg.Headers.Add["X-Proxy-By"])
	}

	if len(cfg.Headers.Remove) != 1 || cfg.Headers.Remove[0] != "X-Powered-By" {
		t.Errorf("Expected header 'X-Powered-By' to be removed, got %q", cfg.Headers.Remove[0])
	}

	if cfg.Timeout != 5000 {
		t.Errorf("Expected Timeout to be 5000, got %d", cfg.Timeout)
	}
}
