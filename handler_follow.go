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

func (cfg *apiConfig) handlerFollowRemove(w http.ResponseWriter, r *http.Request, users database.User) {

	if r.Method != http.MethodDelete {
		respondWIthError(w, http.StatusNotAcceptable, ("Get method not allowed on path"))
	}

	FoId := (r.PathValue("Follow_id"))
	fooId, err := uuid.Parse(FoId)
	if err != nil {
		respondWIthError(w, http.StatusNotFound, "couldnt not delete")
		return

	}
	err = cfg.DB.DeleteFollow(r.Context(), fooId)
	if err != nil {
		clog.Printf("failed to remove follow\n")
		respondWIthError(w, http.StatusNotModified, "failed to remove follow")

	}
	clog.Printf("Feed Removed \n# %s\n", fooId)
	respondWIthError(w, 200, "Feed Deleted")

}

func (cfg *apiConfig) handlerFollowsGET(w http.ResponseWriter, r *http.Request, users database.User) {

	if r.Method != http.MethodGet {
		respondWIthError(w, http.StatusNotAcceptable, ("Get method not allowed on path"))
		return
	}
	folw, err := cfg.DB.GetFollowsAll(r.Context(), users.ID)
	if err != nil {
		respondWIthError(w, http.StatusNotFound, "not found")
	}
	clog.Printf("Sucess\nFetched all follows for user: %v", users.ID)
	respondWithJSON(w, 200, FollowToFollows(folw))

}
