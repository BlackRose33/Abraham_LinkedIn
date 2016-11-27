package bootstrap

import (
	"net/http"

	"abraham_linkedin/controllers"
	"abraham_linkedin/resources"
)

// CreateRoutes maps URI's to the corresponding controller method
func CreateRoutes() {
	resources.MapCSSHandler()
	resources.MapImageHandler()
	resources.MapJSHandler()

	http.Handle(route("/", controllers.Home))
	http.Handle(route("/candinfo/", controllers.CandInfo))
	http.Handle(route("/highspender", controllers.HighSpender))
	http.Handle(route("/mostpaid", controllers.MostPaid))
	http.Handle(route("/trends", controllers.Trends))
	http.Handle(route("/borough/heat/contrib/", controllers.BoroughHeatmapContribCandid))
}

// Quick wrapper for StripPrefix which prevents typos
func route(path string, callback http.HandlerFunc) (string, http.Handler) {
	return path, http.StripPrefix(path, callback)
}
