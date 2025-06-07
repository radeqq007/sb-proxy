package proxy

import (
	"log"
	"net/http"
	"sb-proxy/internal/config"
	"strconv"
	"time"
)

func NewRouter(cfg *config.Config) *http.Server {
	server := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Port),
		ReadTimeout:  time.Duration(cfg.Timeout) * time.Millisecond,
		WriteTimeout: time.Duration(cfg.Timeout) * time.Millisecond,
		IdleTimeout:  time.Duration(cfg.Timeout) * time.Millisecond,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Incoming request: %s %s", r.Method, r.URL)

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
