package main

import (
	"dochub/Controllers"
	"dochub/lib"
	"net/http"
)

func main() {
	lib.ConnectToDB()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", methodHandler(http.MethodGet, Controllers.IndexHandler))
	http.HandleFunc("/login", methodHandler(http.MethodGet, Controllers.LoginHandler))
	http.HandleFunc("/register", methodHandler(http.MethodGet, Controllers.RegisterHandler))
	http.HandleFunc("/forget-password", methodHandler(http.MethodGet, Controllers.ForgetPasswordHandler))

	//http.HandleFunc("/dashboard", dashboardHandler)
	//
	http.HandleFunc("/sign-in", methodHandler(http.MethodPost, Controllers.SignInHandler))
	http.HandleFunc("/sign-up", methodHandler(http.MethodPost, Controllers.SignUpHandler))

	errHTTP := http.ListenAndServe(":3000", nil)
	if errHTTP != nil {
		return
	}

}

//
//func dashboardHandler(w http.ResponseWriter, r *http.Request) {
//	err := tpl.ExecuteTemplate(w, "dashboard.html", nil)
//	if err != nil {
//		return
//	}
//}
