package controllers

import (
	"abraham_linkedin/models"
	"strconv"
	"net/http"
)

// MostPaid handles mostpaid/##candid
func MostPaid(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)
	viewData := BaseViewData(w, r)

	if len(args) != 1 {
		args = append(args, "")
	}

	yearStr := args[0]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	amounts, err := models.GetHighestPaidByPositionForYear(Base.Db, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = amounts

	RenderView(w, "layouts#MostPaid", viewData)
}
