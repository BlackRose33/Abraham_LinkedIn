package controllers

import (
	"abraham_linkedin/models"
	"fmt"
	"net/http"
)

type candInfoViewData struct {
	Candidate    *models.Candidate
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

	candidateInfo, err := models.GetCandidateByID(Base.Db, args[0])

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	data := candInfoViewData{}

	if candidateInfo == nil {
		candidateInfo = &models.Candidate{
			ID:        " ",
			FirstName: " ",
			LastName:  " ",
		}
		data.Candidate = candidateInfo
	} else {

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
		data.Candidate = candidateInfo
		data.Summaries = candAmountSummaries
		data.History = candHistory
		data.Contribs = topContribs
		data.Expenses = topExpenses
		data.Reasons = topReasons
		data.Reasons2 = topReasons2
		data.PrizesGiven = prizesGiven
		data.PrizesBought = prizesBought
		data.CFPData = cfpData
	}

	viewData.Data = &data

	RenderView(w, "layouts#CandInfo", viewData)
}
