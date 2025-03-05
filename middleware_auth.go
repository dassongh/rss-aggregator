package main

import (
	"fmt"
	"net/http"

	"github.com/dassongh/rss-aggregator/internal/auth"
	"github.com/dassongh/rss-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("User not found: %v", err))
			return
		}

		handler(w, r, user)
	}
}
