package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var errors []string

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", methodHandler(http.MethodGet, indexHandler))
	http.HandleFunc("/login", methodHandler(http.MethodGet, loginHandler))
	http.HandleFunc("/register", methodHandler(http.MethodGet, registerHandler))
	http.HandleFunc("/forget-password", methodHandler(http.MethodGet, forgetPasswordHandler))

	//http.HandleFunc("/dashboard", dashboardHandler)
	//
	http.HandleFunc("/sign-in", methodHandler(http.MethodPost, signInHandler))
	http.HandleFunc("/sign-up", methodHandler(http.MethodPost, signUpHandler))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}

func render(w http.ResponseWriter, name string, header bool, data interface{}) {
	funcMap := template.FuncMap{
		"old": old, // register your helper
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//tpl := template.Must(template.ParseFiles("templates/layouts/app.html", "templates/layouts/header.html", "templates/index.html"))

	render(w, "index.html", true, map[string]interface{}{
		"Title": "Home Page",
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "login.html", false, map[string]interface{}{
		"Title": "Login Page",
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "register.html", false, map[string]interface{}{
		"Title": "Register Page",
	})
}

func forgetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "forget-password.html", false, map[string]interface{}{
		"Title": "Forget Password Page",
	})
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "dashboard.html", nil)
	if err != nil {
		return
	}
}
