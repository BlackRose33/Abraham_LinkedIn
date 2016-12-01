package controllers

import (
	"abraham_linkedin/models"
	"net/http"
)

// BoroughHeatmapContribAmtCandid handles the route '/graphs/amt/#cand_id'
func BoroughHeatmapContribAmtCandid(w http.ResponseWriter,
	r *http.Request) {

	args := URIArgs(r)
	viewData := BaseViewData(w, r)

	if len(args) != 1 {
		args = append(args, "")
	}

	boroughCounts, err := models.GetBoroughContribAmountPercentagesForCandidate(
		Base.Db, args[0])

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	viewData.Data = boroughCounts

	RenderView(w, "borough#contribheat", viewData)
}

// BoroughHeatmapContribCandid handles the route
// '/graphs/##candid'
func BoroughHeatmapContribCandid(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)
	viewData := BaseViewData(w, r)

	if len(args) != 1 {
		args = append(args, "")
	}

	boroughCounts, err := models.GetBoroughContribPercentagesForCandidate(Base.Db,
		args[0])

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	viewData.Data = boroughCounts

	RenderView(w, "layouts#Graphs", viewData)
}
