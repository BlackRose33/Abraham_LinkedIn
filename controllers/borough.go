package controllers

import (
	"abraham_linkedin/models"
	"net/http"
)

// BoroughHeatmapContribCandid handles the route
// '/borough/heat/contrib/##candid'
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
