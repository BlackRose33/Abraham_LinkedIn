package bootstrap

import (
	"html/template"
	"strings"
)

// GetTemplates load templates with their corresponding layouts and stores
// them in a map
func GetTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)

	templates["home#index"] = loadTemplate("views/index.html")

	return templates
}

func loadTemplate(fpath string) *template.Template {

	funcs := template.FuncMap{
		"title":   strings.Title,
		"compare": strings.Compare,
	}

	return template.Must(
		template.New("").Funcs(funcs).ParseFiles(
			"views/layouts/default.html",
			fpath))
}
