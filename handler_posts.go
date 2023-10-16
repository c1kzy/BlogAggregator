package main

import (
	"blogAggregator/internal/database"
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := cfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't get posts for user")
		return
	}

	respondwithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
