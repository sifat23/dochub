package main

import (
	"fmt"
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
	http.HandleFunc("/forget-password", forgetPasswordHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}

func render(w http.ResponseWriter, name string, header bool, data interface{}) {
	tpl := template.Must(template.ParseFiles(
		"templates/layouts/app.html",
		"templates/layouts/header.html",
		"templates/"+name,
	))

	fmt.Print(data)

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
