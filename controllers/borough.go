package controllers

import (
	"abraham_linkedin/models"
	"net/http"
)

type boroughHeatmapData struct {
	BoroughPercentages models.BoroughPercentages
	CandID             string
	Candidate          *models.Candidate
}

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

	if err != nil || boroughCounts == nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	var candidate = &models.Candidate{
		ID:        args[0],
		FirstName: "All",
		LastName:  "Candidates",
	}
	if len(args[0]) > 0 {
		candidateTmp, err := models.GetCandidateByID(Base.Db, args[0])
		if candidateTmp != nil && err == nil {
			candidate = candidateTmp
		}
	}

	viewData.Data = &boroughHeatmapData{
		BoroughPercentages: *boroughCounts,
		CandID:             args[0],
		Candidate:          candidate,
	}

	RenderView(w, "layouts#Graphs", viewData)
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

	if err != nil || boroughCounts == nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	var candidate = &models.Candidate{
		ID:        args[0],
		FirstName: "All",
		LastName:  "Candidates",
	}
	if len(args[0]) > 0 {
		candidateTmp, err := models.GetCandidateByID(Base.Db, args[0])
		if candidateTmp != nil && err == nil {
			candidate = candidateTmp
		}
	}

	viewData.Data = &boroughHeatmapData{
		BoroughPercentages: *boroughCounts,
		CandID:             args[0],
		Candidate:          candidate,
	}

	RenderView(w, "layouts#Graphs", viewData)
}
