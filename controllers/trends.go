package controllers

import (
	"net/http"
)

// Trends handles trends/##candid
func Trends(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)
	viewData := BaseViewData(w, r)

	if len(args) != 1 {
		args = append(args, "")
	}

	RenderView(w, "layouts#Trends", viewData)
}
