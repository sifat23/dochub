package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/layouts/app.html", "templates/index.html"))

	data := map[string]interface{}{
		"Title": "Home Page",
	}
	err := tpl.ExecuteTemplate(w, "app.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/layouts/app.html", "templates/login.html"))

	err := tpl.ExecuteTemplate(w, "app.html", nil)
	if err != nil {
		return
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/layouts/app.html", "templates/register.html"))

	err := tpl.ExecuteTemplate(w, "app.html", nil)
	if err != nil {
		return
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "dashboard.html", nil)
	if err != nil {
		return
	}
}
