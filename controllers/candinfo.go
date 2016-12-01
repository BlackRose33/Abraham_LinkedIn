package controllers

import (
	"abraham_linkedin/models"
	"net/http"
)

// CandInfo handles candinfo/##candid
func CandInfo(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)
	viewData := BaseViewData(w, r)

	if len(args) != 1 {
		args = append(args, "")
	}

	CandidateInfo, err := models.GetCandidateByIDWithInfo(Base.Db,
		args[0])

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	viewData.Data = CandidateInfo

	RenderView(w, "layouts#CandInfo", viewData)
}
