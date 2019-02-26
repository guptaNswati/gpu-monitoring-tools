package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// res: curl localhost:9400/gpu/metrics

func main() {
	stopSig := make(chan os.Signal, 1)
	signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM)

	addr := ":9400"
	server := newHttpServer(addr)

	go func() {
		log.Printf("Running http server on localhost%s", addr)
		server.serve()
	}()
	defer server.stop()

	<-stopSig
	return
}
