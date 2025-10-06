package Controllers

import (
	"dochub/lib"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, "index.html", true, map[string]interface{}{
		"Title": "Home Page",
	})
}
