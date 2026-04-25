package main

import (
	"api/controller"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	socketPath := os.Getenv("RINHA_SOCKET_PATH")
	if socketPath == "" {
		socketPath = "/tmp/api.sock"
	}

	if err := os.RemoveAll(socketPath); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /ready", controller.ReadyHandle)
	http.HandleFunc("POST /fraud-score", controller.FraudScoreHandle)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal("Erro ao abrir socket unix:", err)
	}

	if err := os.Chmod(socketPath, 0777); err != nil {
		log.Fatal("Erro nas permissões do socket:", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		os.Remove(socketPath)
		os.Exit(0)
	}()

	log.Printf("API a ouvir no socket unix: %s", socketPath)	
	log.Fatal(http.Serve(listener, nil))
}
