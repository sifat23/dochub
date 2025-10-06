package routes

import (
	"dochub/Controllers"
	"dochub/bin"
	"net/http"
	"strings"
)

func methodHandler(method string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(bin.Errors) > 0 {
			http.Error(w, strings.Join(bin.Errors, ", "), http.StatusBadRequest)
			return
		}

		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}

func Routes() {
	http.HandleFunc("/", methodHandler(http.MethodGet, Controllers.IndexHandler))
	http.HandleFunc("/login", methodHandler(http.MethodGet, Controllers.LoginHandler))
	http.HandleFunc("/register", methodHandler(http.MethodGet, Controllers.RegisterHandler))
	http.HandleFunc("/forget-password", methodHandler(http.MethodGet, Controllers.ForgetPasswordHandler))

	//http.HandleFunc("/dashboard", dashboardHandler)
	//
	http.HandleFunc("/sign-in", methodHandler(http.MethodPost, Controllers.SignInHandler))
	http.HandleFunc("/sign-up", methodHandler(http.MethodPost, Controllers.SignUpHandler))
}
