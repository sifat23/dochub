package main

import (
	"dochub/bin/routes"
	"dochub/lib"
	"net/http"
)

func main() {
	lib.ConnectToDB()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	routes.Routes()

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
