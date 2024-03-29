package main

import (
	"blogAggregator/internal/database"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	respondwithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (cfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondwithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

func (cfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondwithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	respondwithJSON(w, http.StatusOK, struct{}{})
}
