package controllers

import (
	"abraham_linkedin/models"
	"fmt"
	"net/http"
)

type candInfoViewData struct {
	Candidate *models.CandInfo
	Summaries []models.CandidateAmountSummary
}

// CandInfo handles candinfo/##candid
func CandInfo(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)
	viewData := BaseViewData(w, r)

	if len(args) != 1 {
		args = append(args, "")
	}

	candidateInfo, err := models.GetCandidateByIDWithInfo(Base.Db,
		args[0])

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	candAmountSummaries, err := models.GetSummariesForCandidate(Base.Db, args[0])
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	viewData.Data = &candInfoViewData{
		Candidate: candidateInfo,
		Summaries: candAmountSummaries,
	}

	RenderView(w, "layouts#CandInfo", viewData)
}
