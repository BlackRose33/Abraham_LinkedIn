package controllers

import (
	"abraham_linkedin/models"
	"net/http"
)

// APICandidates handles the route '/api/candidates/'
func APICandidates(w http.ResponseWriter, r *http.Request) {
	candidates, err := models.GetCandidates(Base.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RenderJSON(w, candidates)
}
