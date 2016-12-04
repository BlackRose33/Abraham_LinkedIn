package controllers

import (
	"net/http"
)

var trendFunctions map[string]http.Handler

func initTrendFunctions() {
	trendFunctions = map[string]http.Handler{
		"": 						http.HandlerFunc(TrendsIndex),
		"ExpChange": 		http.HandlerFunc(ExpChange),
		"ConChange": 		http.HandlerFunc(ConChange),
		"HighSpender": 	http.HandlerFunc(HighestSpender),
		"HighContrib": 	http.HandlerFunc(HighestContrib),
		"MostPaid": 		http.HandlerFunc(Mostpaid),
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

	RenderView(w, "trends#ExpChange", viewData)
}

//ConChange handles /trends/#queryName
func ConChange(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	RenderView(w, "trends#ConChange", viewData)
}

//HighSpender handles /trends/#queryName
func HighestSpender(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	RenderView(w, "trends#HighSpender", viewData)
}

//HighContrib handles /trends/#queryName
func HighestContrib(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	RenderView(w, "trends#HighContrib", viewData)
}

//MostPaid handles /trends/#queryName
func Mostpaid(w http.ResponseWriter, r *http.Request) {
	viewData := BaseViewData(w, r)

	RenderView(w, "trends#MostPaid", viewData)
}
