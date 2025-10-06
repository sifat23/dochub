package Controllers

import (
	"dochub/lib"
	"net/http"
)

func ForgetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, "forget-password.html", false, map[string]interface{}{
		"Title": "Forget Password Page",
	})
}
