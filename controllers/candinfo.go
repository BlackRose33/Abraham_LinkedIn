package controllers

import (
	"abraham_linkedin/models"
	"fmt"
	"net/http"
)

type candInfoViewData struct {
	Candidate    *models.CandInfo
	Summaries    []models.CandidateAmountSummary
	History      []models.Candidacy
	Contribs     []models.Contributor
	Expenses     []models.Expense
	Reasons      []models.Reason
	Reasons2     []models.Reason
	PrizesGiven  []models.DoorPrize
	PrizesBought []models.DoorPrize
	CFPData      *models.CandidateCFPData
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

	candHistory, err := models.GetCandidacyHistory(Base.Db, args[0])
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	topExpenses, err := models.GetTopTenExpensesForCandidate(Base.Db, args[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	topContribs, err := models.GetTopTenContributorsForCandidate(Base.Db, args[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	topReasons, err := models.GetTopReasonsForCandidate(Base.Db, args[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	topReasons2, err := models.GetPriceyReasonsForCandidate(Base.Db, args[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prizesGiven, err := models.GetTopDoorPrizesGiven(Base.Db, args[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prizesBought, err := models.GetTopDoorPrizesPurchased(Base.Db, args[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cfpData, err := models.GetCandCFPData(Base.Db, args[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = &candInfoViewData{
		Candidate:    candidateInfo,
		Summaries:    candAmountSummaries,
		History:      candHistory,
		Contribs:     topContribs,
		Expenses:     topExpenses,
		Reasons:      topReasons,
		Reasons2:     topReasons2,
		PrizesGiven:  prizesGiven,
		PrizesBought: prizesBought,
		CFPData:      cfpData,
	}

	RenderView(w, "layouts#CandInfo", viewData)
}
