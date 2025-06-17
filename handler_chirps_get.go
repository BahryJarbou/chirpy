package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, req *http.Request) {
	dbchirps, err := cfg.db.GetChrips(req.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	chirps := []Chirp{}
	for _, chirp := range dbchirps {
		chirps = append(chirps, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	respondWithJson(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpGet(w http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(req.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "invalid ID", err)
		return
	}

	dbchirp, err := cfg.db.GetChirpByID(req.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not retrieve chirp", err)
		return
	}

	respondWithJson(w, http.StatusOK, Chirp{
		ID:        dbchirp.ID,
		CreatedAt: dbchirp.CreatedAt,
		UpdatedAt: dbchirp.UpdatedAt,
		Body:      dbchirp.Body,
		UserID:    dbchirp.UserID,
	})
}
