package main

import (
	"encoding/json"
	"github.com/google/uuid"
	//	"github.com/gpr3211/blogger/internal/auth"
	"github.com/gpr3211/blogger/internal/clog"
	"github.com/gpr3211/blogger/internal/database"
	"net/http"
	"time"
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

	respondWithJSON(w, http.StatusOK, dbToUser(user))
}

/*
func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	head, err := auth.GetApiHead(r.Header)
	if err != nil {
		return
	}
//	token := cfg.DB.GetUserByAPIKey(r.Context(), head)
}





*/
