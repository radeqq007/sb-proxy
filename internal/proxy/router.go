package proxy

import (
	"net/http"
	"sb-proxy/internal/config"
	"strconv"
)

func NewRouter(cfg *config.Config) *http.Server {
	server := &http.Server{
		Addr: ":" + strconv.Itoa(cfg.Port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Add headers
			for key, value := range cfg.Headers.Add {
				w.Header().Set(key, value)
			}

			// Remove headers
			for _, key := range cfg.Headers.Remove {
				w.Header().Del(key)
			}

			for _, route := range cfg.Routes {
				if len(r.URL.Path) >= len(route.PathPrefix) && r.URL.Path[:len(route.PathPrefix)] == route.PathPrefix {
					proxy := newProxy(route.Target)
					proxy.ServeHTTP(w, r)
					return
				}
			}
			http.NotFound(w, r)
		}),
	}

	return server
}
