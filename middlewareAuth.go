package main

import (
	"net/http"

	"github.com/gpr3211/blogger/internal/auth"
	"github.com/gpr3211/blogger/internal/clog"
	"github.com/gpr3211/blogger/internal/database"
)

// CUSTOM TYPE FOR HANDLERS THAT REQ AUTH

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		head, err := auth.GetApiHead(r.Header)
		if err != nil {
			clog.Printf("Failed to get key %s\n", head)
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), head)
		if err != nil {
			clog.Printf("Failed to fetch user Api auth\n")
			respondWIthError(w, http.StatusUnauthorized, "failed to fetch user from db")
		}
		handler(w, r, user)
	}
}
