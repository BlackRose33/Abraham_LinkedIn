package controllers

import (
	"net/http"
)

// CandInfo handles candinfo/##candid
func CandInfo(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)
	viewData := BaseViewData(w, r)

	if len(args) != 1 {
		args = append(args, "")
	}

	RenderView(w, "layouts#CandInfo", viewData)
}
