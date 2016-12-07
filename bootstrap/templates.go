package bootstrap

import (
	"html/template"
	"strconv"
	"strings"
)

// GetTemplates load templates with their corresponding layouts and stores
// them in a map
func GetTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)

	// Main pages
	templates["home#index"] = loadTemplate("views/index.html")
	templates["layouts#CandInfo"] = loadTemplate("views/layouts/CandInfo.html")
	templates["layouts#HighSpender"] = loadTemplate("views/layouts/HighSpender.html")
	templates["layouts#MostPaid"] = loadTemplate("views/layouts/MostPaid.html")
	templates["layouts#Graphs"] = loadTemplate("views/layouts/Graphs.html")

	// Individual trends
	templates["trends#index"] = loadTrendsTemplate("views/trends/index.html")
	templates["trends#HighSpender"] = loadTrendsTemplate("views/trends/HighSpender.html")
	templates["trends#HighContrib"] = loadTrendsTemplate("views/trends/HighContrib.html")
	templates["trends#MostPaid"] = loadTrendsTemplate("views/trends/MostPaid.html")
	templates["trends#ExpChange"] = loadTrendsTemplate("views/trends/ExpChange.html")
	templates["trends#ConChange"] = loadTrendsTemplate("views/trends/ConChange.html")
	templates["trends#explanations"] = loadTrendsTemplate("views/trends/explanations.html")
	templates["trends#amtExpln"] = loadTrendsTemplate("views/trends/amtExpl.html")
	templates["trends#refunds"] = loadTrendsTemplate("views/trends/refunds.html")
	templates["trends#cfptrend"] = loadTrendsTemplate("views/trends/cfptrend.html")
	templates["trends#avgcfptrend"] = loadTrendsTemplate("views/trends/avgcfptrend.html")
	templates["trends#occupations"] = loadTrendsTemplate("views/trends/occupations.html")
	templates["trends#occupations2"] = loadTrendsTemplate("views/trends/occupations2.html")

	return templates
}

func formatPrice(price float64) string {
	return "$" + strconv.FormatFloat(price, 'f', 2, 64)
}

func loadTrendsTemplate(fpath string) *template.Template {

	funcs := template.FuncMap{
		"title":   strings.Title,
		"compare": strings.Compare,
		"price":   formatPrice,
	}

	return template.Must(
		template.New("").Funcs(funcs).ParseFiles(
			"views/layouts/default.html",
			"views/layouts/Trends.html",
			fpath))

}

func loadTemplate(fpath string) *template.Template {

	funcs := template.FuncMap{
		"title":   strings.Title,
		"compare": strings.Compare,
		"price":   formatPrice,
	}

	return template.Must(
		template.New("").Funcs(funcs).ParseFiles(
			"views/layouts/default.html",
			fpath))
}
