package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gpr3211/blogger/internal/clog"
	"github.com/gpr3211/blogger/internal/database"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	//	METRICS
	mux.HandleFunc("GET /blog/v1/healthz", cfg.handleReady)
	mux.HandleFunc("GET /blog/v1/err", cfg.handleError)
	// USERS
	mux.HandleFunc("POST /blog/v1/users", cfg.handlerCreateUser)
	mux.HandleFunc("GET /blog/v1/users", cfg.middlewareAuth(cfg.handlerGetUser))
	// FEEDS

	mux.HandleFunc("POST /blog/v1/feeds", cfg.middlewareAuth(cfg.handlerFeedCreate))
	mux.HandleFunc("GET /blog/v1/feeds", cfg.handlerFeedsGet)

	// FOLLOWS
	mux.HandleFunc("POST /blog/v1/feed_follows", cfg.middlewareAuth(cfg.handlerFollowCreate))
	mux.HandleFunc("DELETE /blog/v1/feed_follows/{Follow_id}", cfg.middlewareAuth(cfg.handlerFollowRemove))
	mux.HandleFunc("GET /blog/v1/feed_follows", cfg.middlewareAuth(cfg.handlerFollowsGET))
	srv := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	clog.Printf("Serving files from %s on port: %s", filepathRoot, port)
	fmt.Println("Starting server..")
	fmt.Printf("Serving files on port %s", port)
	dat_ass := database.New(db)
	go startScrapeWorker(dat_ass, 10, time.Minute)

	log.Fatal(srv.ListenAndServe())

}
