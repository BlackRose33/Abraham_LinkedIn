package controllers

import (
	"net/http"
)

//mostpaid/##candid
func MostPaid(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)
	viewData := BaseViewData(w, r)

	if len(args) != 1 {
		args = append(args, "")
	}

	RenderView(w, "layouts#MostPaid", viewData)
}
