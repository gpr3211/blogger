package main

import (
	"net/http"
	"strconv"

	"github.com/gpr3211/blogger/internal/clog"
	"github.com/gpr3211/blogger/internal/database"
)

func (cfg *apiConfig) handlerGetPostsUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limit := (r.PathValue("limit"))
	l, err := strconv.Atoi(limit)
	if err != nil {
		respondWIthError(w, http.StatusNoContent, "Must provide limit query param")
	}

	posts, err := cfg.DB.GetUserPosts(r.Context(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  int32(l),
	})
	if err != nil {
		clog.Printf("Failed to fetch posts from database for User: %v", user.ID)
	}
	output := dbToPosttoPosts(posts)
	clog.Printf("GET v1/posts Posts Successfully Fetched for User: %v", user.ID)
	respondWithJSON(w, 200, output)
}
