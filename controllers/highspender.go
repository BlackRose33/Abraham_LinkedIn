package controllers

import (
	"abraham_linkedin/models"
	"net/http"
	"strconv"
)

// HighSpender handles highspender/##candid
func HighSpender(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)
	viewData := BaseViewData(w, r)

	if len(args) != 1 {
		args = append(args, "2001")
	}

	yearStr := args[0]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	amounts, err := models.GetHighestSpenderByPositionForYear(Base.Db, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = amounts

	RenderView(w, "layouts#HighSpender", viewData)
}
