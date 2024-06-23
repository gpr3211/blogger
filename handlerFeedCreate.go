package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	//	"github.com/gpr3211/blogger/internal/auth"
	"github.com/gpr3211/blogger/internal/clog"
	"github.com/gpr3211/blogger/internal/database"
)

func (cfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	if r.Method != http.MethodPost {
		respondWIthError(w, http.StatusNotAcceptable, ("Get method not allowed on path"))
	}

	type parameters struct {
		Name string
		URL  string
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		clog.C.Printf("error decoding json body")
	}
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	clog.Printf("Successfully created Feed for User: %v\nName: %s\nUrl: %s\n", feed.UserID, feed.Name, feed.Url)
	respondWithJSON(w, 200, dbToFeed(feed))

}
