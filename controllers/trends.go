package controllers

import (
	"net/http"
)

var trendFunctions map[string]http.Handler

func initTrendFunctions() {
	trendFunctions = map[string]http.Handler{
		"": http.HandlerFunc(TrendsIndex),
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
