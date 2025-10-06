package Controllers

import (
	"dochub/bin/models"
	"dochub/lib"
	"fmt"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	msg := &models.Message{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if msg.SignInFormValidate() == false {
		lib.Render(w, "login.html", false, map[string]interface{}{
			"Title": "Home Page",
			"msg":   msg,
		})
		return
	}

	// Process the data (e.g., save to database, send email)
	fmt.Fprintf(w, "Thank you, %s! Your email %s has been received.", msg.Name, msg.Password)
}
