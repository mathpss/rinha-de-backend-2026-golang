package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

var backends = []string{
	"/sockets/api-1.sock",
	"/sockets/api-2.sock",
}

var counter uint64

func main() {
	transport := &http.Transport{
		
		DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
			
			idx := atomic.AddUint64(&counter, 1) % uint64(len(backends))

			
			conn, err := net.Dial("unix", backends[idx])
			if err != nil {
			
				nextIdx := (idx + 1) % uint64(len(backends))
				log.Printf("Backend %s offline, a tentar %s", backends[idx], backends[nextIdx])
				return net.Dial("unix", backends[nextIdx])
			}
			return conn, nil
		},
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		MaxIdleConnsPerHost: 64, 
	}

	server := &http.Server{
		Addr:    ":9999",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			serveProxy(w, r, transport)
		}),
	}

	log.Println("Load Balancer Go iniciado em :9999 (via Sockets)")
	log.Fatal(server.ListenAndServe())
}

func serveProxy(w http.ResponseWriter, r *http.Request, transport *http.Transport) {
	
	outReq := r.Clone(r.Context())
	
	outReq.URL.Scheme = "http"
	outReq.URL.Host = "unix-provider"

	resp, err := transport.RoundTrip(outReq)
	if err != nil {
		http.Error(w, "Serviço Indisponível", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}