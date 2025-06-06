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
}
*/

type Route struct {
	PathPrefix string `json:"path_prefix"`
	Target     string `json:"target"`
}

type Config struct {
	Port   int     `json:"port"`
	Routes []Route `json:"routes"`
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
