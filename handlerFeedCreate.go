package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	//	"github.com/google/uuid"
	//	"github.com/gpr3211/blogger/internal/auth"
	"github.com/google/uuid"
	"github.com/gpr3211/blogger/internal/clog"
	"github.com/gpr3211/blogger/internal/database"
)

func (cfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	if r.Method != http.MethodPost {
		respondWIthError(w, http.StatusNotAcceptable, ("Get method not allowed on path"))
		return
	}

	type parameters struct {
		Name string
		URL  string
	}
	type response struct {
		Feed   Feed   `json:"feed"`
		Follow Follow `json:"feed_follow"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		clog.C.Printf("error decoding json body")
	}
	fuuid := uuid.New()

	// FEED CREATE
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        fuuid,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		LastFetch: sql.NullTime{},
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		clog.Printf("failed to create feed %v", err)
		respondWIthError(w, http.StatusNotImplemented, "Feed not created")
		return
	}
	// AUTO FOLLOW BY USER
	follow, err := cfg.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    fuuid,
	})

	if err != nil {
		clog.Printf("follow not created for user %v\n", follow.UserID)
		respondWIthError(w, http.StatusNotModified, "Follow not created")

	}
	res := response{
		Feed:   dbToFeed(feed),
		Follow: dbToFollow(follow),
	}

	clog.Printf("Successfully created Feed for User: %v\nName: %s\nUrl: %s\n", feed.UserID, feed.Name, feed.Url)
	respondWithJSON(w, 200, res)

}

func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWIthError(w, http.StatusNotAcceptable, ("Get method not allowed on path"))
	}
	datt, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWIthError(w, 404, "nah")
	}
	respondWithJSON(w, 200, FeedToFeeds(datt))
}
