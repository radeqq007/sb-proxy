package main

import (
	"sb-proxy/internal/config"
	"sb-proxy/internal/proxy"
)

func main() {
	cfgPath := "sb-config.json"
	cfg, err := config.Load(cfgPath)
	if err != nil {
		panic(err)
	}

	server := proxy.NewRouter(cfg)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	defer server.Close()
}
