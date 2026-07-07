package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ZooHandler struct {
}

func NewZooHandler() *ZooHandler {
	return &ZooHandler{}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if body != nil {
		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Printf("failed to encode response: %v", err)
		}
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// GET /api/v1/zoos/{enclosures}
func (h *ZooHandler) Get(w http.ResponseWriter, r *http.Request) {
	enclosures := r.PathValue("enclosures")

	n, err := strconv.Atoi(enclosures)
	if err != nil {
		writeError(w, http.StatusBadRequest , "Invalid Number")
		return
	}

	if n <= 0 {
		writeError(w, http.StatusBadRequest , "Number of enclosures needs to be > 0")
		return
	}

	if n > 100 {
		writeError(w, http.StatusBadRequest, "number of enclosures exceeds 100")
		return
	}

	animals, animalsErr := LoadAnimals()

	if animalsErr != nil {
		log.Printf("LoadAnimals failed: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to load animal data")
		return
	}

	zoo, algoErr := ComputeZoo(n, animals)

	if algoErr != nil {
		log.Printf("ComputeZoo failed for n=%d: %v", n, err)
		writeError(w, http.StatusInternalServerError, "failed to compute zoo layout")
		return
	}

	writeJSON(w, http.StatusOK, zoo)
}
