package main

import (
	//	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gpr3211/blogger/internal/auth"
	"github.com/gpr3211/blogger/internal/clog"
	"github.com/gpr3211/blogger/internal/database"
)

type parameters struct {
	Name string
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWIthError(w, http.StatusNotAcceptable, ("Get method not allowed on path"))
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		clog.C.Printf("error decoding json body")
	}
	user, err := cfg.DB.CrateUser(r.Context(), database.CrateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	clog.Printf("Created User: %v\n Name: %s \n", user.ID, user.Name)
	respondWithJSON(w, http.StatusOK, dbToUser(user))
}

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	head, err := auth.GetApiHead(r.Header)
	if err != nil {
		return
	}
	user, err := cfg.DB.GetUserByAPIKey(r.Context(), head)
	if err != nil {
		clog.Printf("Failed to get user by ApiKey")
		respondWIthError(w, http.StatusUnauthorized, "failed to auth api")
	}
	clog.Printf("Getting user info successful by ID:%s", user.ID)
	respondWithJSON(w, http.StatusFound, user)

}
