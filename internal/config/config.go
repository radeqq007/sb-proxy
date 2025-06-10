package config

import (
	"encoding/json"
	"os"
)

/*
example sb-config.json:
{
	"port": 8080,
	"routes": [
		{
      "path_prefix": "/api/",
      "target": "http://localhost:3000"
    },
    {
      "path_prefix": "/admin/",
      "target": "http://localhost:4000"
    },
	],
	"headers": {
		"add": {
			"X-Proxy-By": "sb-proxy",
		},
		"remove": ["X-Powered-By"]
	},
	"timeout_ms": 30,
  "rate_limit": {
    "enabled": true,
    "requests_per_minute": 60,
    "burst": 10,
    "cleanup_interval_minutes": 5,
  }
}
*/

type Route struct {
	PathPrefix string `json:"path_prefix"`
	Target     string `json:"target"`
}

type Config struct {
	Port    int     `json:"port"`
	Routes  []Route `json:"routes"`
	Timeout int     `json:"timeout_ms"` // in milliseconds
	Headers struct {
		Add    map[string]string `json:"add"`
		Remove []string          `json:"remove"`
	} `json:"headers"`
	RateLimit struct {
		Enabled           bool `json:enabled`
		RequestsPerMinute int  `json:requests_per_minute`
		Burst             int  `json:burst`
		cleanupInterval   int  `json:cleanup_interval_minutes`
	} `json: "rate_limit`
}

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
