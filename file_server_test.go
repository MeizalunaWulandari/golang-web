package golang_web

import (
	// "fmt"
	"testing"
	"net/http"
)

func TestFileServer404(t *testing.T){
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", fileServer)

	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	// HASIL 404 NOT FOUND
}

func TestFileServer(t *testing.T){
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}