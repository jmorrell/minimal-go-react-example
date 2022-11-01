package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	server http.Server
}

func endpointHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello\n")
}

func endpointWithIDHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello: %s\n", id)
}

func NewServer() server {
	r := mux.NewRouter()
	r.HandleFunc("/endpoint", endpointHandler)
	r.HandleFunc("/endpoint/{id}", endpointWithIDHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return server{
		server: http.Server{
			Addr:    "localhost:8000",
			Handler: r,
		},
	}
}

func (s *server) Start() error {
	return s.server.ListenAndServe()
}

func (s *server) Close() error {
	return s.server.Close()
}

func main() {
	server := NewServer()
	err := server.Start()
	defer server.Close()
	if err != nil && err != http.ErrAbortHandler {
		fmt.Println(err.Error())
	}
}
