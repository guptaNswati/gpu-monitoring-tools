package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	timeout    = 5 * time.Second
	gpuMetrics = "/run/dcgm/dcgm-pod.prom"
)

type httpServer struct {
	router *mux.Router
	server *http.Server
}

func newHttpServer(addr string) *httpServer {
	r := mux.NewRouter()

	s := &httpServer{
		router: r,
		server: &http.Server{
			Addr:         addr,
			Handler:      r,
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
		},
	}

	path := "/gpu/metrics"
	s.router.HandleFunc(path, getGPUmetrics).Methods("GET")
	return s
}

func (s *httpServer) serve() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Error: %v", err)
	}
}

func (s *httpServer) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Println("http server stopped")
	}
}

func getGPUmetrics(resp http.ResponseWriter, req *http.Request) {
	metrics, err := ioutil.ReadFile(gpuMetrics)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		log.Printf("error: %v%v: %v", req.Host, req.URL, err.Error())
		return
	}
	resp.Write(metrics)
}
