package controllers

import (
	"net/http"
	"strings"
)

// Home is the index page
func Home(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)

	// This is a catch-all handler, so we need to avoid the double
	// "favicon.ico" request
	if len(args) > 0 && strings.Compare(args[0], "favicon.ico") == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	RenderView(w, "home#index", BaseViewData(w, r))
}
