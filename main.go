package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tpl, _ = tpl.ParseGlob("templates/*.html")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		return
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "register.html", nil)
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
