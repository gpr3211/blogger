package main

import (
	"database/sql"
	"github.com/gpr3211/blogger/internal/clog"
	"github.com/gpr3211/blogger/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB             *database.Queries
	fileServerHits int
	Secret         string
}

func main() {
	const filepathRoot = "."
	// LOAD envi
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env")

	}
	port := (os.Getenv("PORT"))
	dbUrl := (os.Getenv("CONN_STRING"))
	//
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		clog.Printf("oops\n")
	}
	dbQueries := database.New(db)

	cfg := apiConfig{
		DB:             dbQueries,
		fileServerHits: 0,
	}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /blog/v1/users", cfg.handlerCreateUser)
	//mux.HandleFunc("GET /blog/v1/users")

	mux.HandleFunc("GET /blog/v1/healthz", cfg.handleReady)
	mux.HandleFunc("GET /blog/v1/err", cfg.handleError)
	srv := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %s", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())

}
