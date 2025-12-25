package main

import (
	"fmt"
	"net/http"

	"github.com/ab01fazl1/scraper/internal/database"
	"github.com/ab01fazl1/scraper/internal/database/auth"
)

type authedHandler func (w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api_key, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("couldn't get apikey: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), api_key)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("no user found: %v", err))
			return
		}
		handler(w, r, user)
	}
}