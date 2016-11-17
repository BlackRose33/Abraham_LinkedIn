package controllers

import (
	"abraham_linkedin/models"
	"net/http"
	"strings"
)

type homeData struct {
	Candidate *models.Candidate
}

// Home is the index page for the admin panel
func Home(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)

	// This is a catch-all handler, so we need to avoid the double
	// "favicon.ico" request
	if len(args) > 0 && strings.Compare(args[0], "favicon.ico") == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	id := "605"
	if len(args) > 0 {
		id = args[0]
	}

	candidate, err := models.GetCandidateByID(Base.Db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData := BaseViewData(w, r)
	viewData.Data = &homeData{
		Candidate: candidate,
	}
	RenderView(w, "home#index", viewData)
}
