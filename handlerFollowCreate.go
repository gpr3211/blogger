package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gpr3211/blogger/internal/clog"
	"github.com/gpr3211/blogger/internal/database"
)

func (cfg *apiConfig) handlerFollowCreate(w http.ResponseWriter, r *http.Request, users database.User) {

	if r.Method != http.MethodPost {
		respondWIthError(w, http.StatusNotAcceptable, ("Get method not allowed on path"))
	}

	type parameters struct {
		Feed_id uuid.UUID
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		clog.C.Printf("error decoding json body")
	}
	follow, err := cfg.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    users.ID,
		FeedID:    params.Feed_id,
	})
	if err != nil {
		clog.Printf("follow not created for user %v\n", follow.UserID)
		respondWIthError(w, http.StatusNotModified, "Follow not created")

	}
	respondWithJSON(w, 200, dbToFollow(follow))
	clog.Printf("Follow successfuly created for User: %s\n Follow: %s", follow.UserID, params.Feed_id)
}
