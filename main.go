package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	fileServerHits int
	Secret         string
}

func main() {
	const filepathRoot = "."

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env")
	}
	port := (os.Getenv("PORT"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /blog/v1/healthz", handleReady)
	mux.HandleFunc("GET /blog/v1/err", handleError)
	srv := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %s", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())

}
