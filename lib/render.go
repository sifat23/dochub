package lib

import (
	"dochub/bin"
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, name string, header bool, data interface{}) {
	funcMap := template.FuncMap{
		"old": bin.Old, // register your helper
	}

	tpl := template.Must(
		template.New("").Funcs(funcMap).ParseFiles(
			"templates/layouts/app.html",
			"templates/layouts/header.html",
			"templates/"+name,
		))

	content := struct {
		Header bool
		Data   interface{}
	}{
		Header: header,
		Data:   data,
	}

	err := tpl.ExecuteTemplate(w, "app.html", content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
