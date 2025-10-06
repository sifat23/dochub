package Controllers

import (
	"dochub/lib"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, "login.html", false, map[string]interface{}{
		"Title": "Login Page",
	})
}
