package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ab01fazl1/scraper/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	// parse the request
	type CreateFeedRequest struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := CreateFeedRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CretedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, err.Error())
	}
	respondWithJSON(w, 201, feed)
}

func (apiCfg *apiConfig) GetFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type GetFeedRequest struct {
		Id uuid.UUID `json:"id"`
	}
	params := GetFeedRequest{}
	// create type > decode to type > get uuid > query based on uuid > return feed as the payload of respond with json
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}
	feed, err := apiCfg.DB.GetFeed(r.Context(), params.Id)
	if err != nil {
		respondWithError(w, 400, err.Error())
	}
	respondWithJSON(w, 200, feed)
}