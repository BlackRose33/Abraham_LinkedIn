package controllers

import "net/http"

// Home is the index page for the admin panel
func Home(w http.ResponseWriter, r *http.Request) {
	RenderView(w, "home#index", BaseViewData(w, r))
}
