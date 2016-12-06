package controllers

import (
	"abraham_linkedin/models"
	"net/http"
)

var trendFunctions map[string]http.Handler

func initTrendFunctions() {
	trendFunctions = map[string]http.Handler{
		"":             http.HandlerFunc(TrendsIndex),
		"ExpChange":    http.HandlerFunc(ExpChange),
		"ConChange":    http.HandlerFunc(ConChange),
		"HighSpender":  http.HandlerFunc(HighestSpender),
		"HighContrib":  http.HandlerFunc(HighestContrib),
		"MostPaid":     http.HandlerFunc(Mostpaid),
		"explanations": http.HandlerFunc(Explanations),
		"amtExpln":     http.HandlerFunc(AmtExpln),
		"refunds":      http.HandlerFunc(Refunds),
	}
}

// Trends handles trends/##candid
func Trends(w http.ResponseWriter, r *http.Request) {
	args := URIArgs(r)
	if len(args) == 0 {
		args = []string{""}
	}
	fn, ok := trendFunctions[args[0]]
	if ok {
		fn.ServeHTTP(w, r)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

// TrendsIndex handles /trends/#queryName
func TrendsIndex(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	RenderView(w, "trends#index", viewData)
}

//ExpChange handles /trends/#queryName
func ExpChange(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	amounts, err := models.GetExpenditureChange(Base.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = amounts

	RenderView(w, "trends#ExpChange", viewData)
}

//ConChange handles /trends/#queryName
func ConChange(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	amounts, err := models.GetContributionChange(Base.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = amounts

	RenderView(w, "trends#ConChange", viewData)
}

//HighestSpender handles /trends/#queryName
func HighestSpender(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	amounts, err := models.GetBiggestSpender(Base.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = amounts

	RenderView(w, "trends#HighSpender", viewData)
}

//HighestContrib handles /trends/#queryName
func HighestContrib(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	amounts, err := models.GetBiggestContributor(Base.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = amounts

	RenderView(w, "trends#HighContrib", viewData)
}

//Mostpaid handles /trends/#queryName
func Mostpaid(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	amounts, err := models.GetMostPaid(Base.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = amounts

	RenderView(w, "trends#MostPaid", viewData)
}

// Explanations handles trends/explanations
func Explanations(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	explanations, err := models.GetExplanationCounts(Base.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = explanations
	RenderView(w, "trends#explanations", viewData)
}

type refundViewData struct {
	ByAmount      []models.Refund
	ByTimeElapsed []models.Refund
}

// Refunds handles trends/refunds
func Refunds(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	byamount, err := models.GetBiggestRefunds(Base.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytimeelapsed, err := models.GetMostDelayedRefunds(Base.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = &refundViewData{
		ByAmount:      byamount,
		ByTimeElapsed: bytimeelapsed,
	}
	RenderView(w, "trends#refunds", viewData)
}

// AmtExpln asdfadsf
func AmtExpln(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	counts, err := models.GetExpensiveExplanations(Base.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	viewData.Data = counts
	RenderView(w, "trends#amtExpln", viewData)
}
