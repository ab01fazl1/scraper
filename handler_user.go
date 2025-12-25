package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ab01fazl1/scraper/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// parse request body
	type parameters struct {
		Name string `json:"name"`
	}
	// what do we get as request body? => json => decode json
	decoder := json.NewDecoder(r.Body)

	// what do we want now? => we want to get the parameters from the body
	// note that
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	// call the query function to create a new user
	user, err := apiCfg.DB.CreatUser(r.Context(), database.CreatUserParams{
		ID:        uuid.New(),
		CretedAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't create user: %v", err))
		return
	}

	// here we are returning the user object sent by the DB, which is not a good practice, 
	// and you should create your own user struct and send it down her
	respondWithJSON(w, 201, user)
}


func (apiCfg *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, user)
}