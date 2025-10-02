package main

import (
	"fmt"
	"net/http"
)

func signInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	msg := &Message{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if msg.signInFormValidate() == false {
		render(w, "login.html", false, map[string]interface{}{
			"Title": "Home Page",
			"msg":   msg,
		})
		return
	}

	// Process the data (e.g., save to database, send email)
	fmt.Fprintf(w, "Thank you, %s! Your email %s has been received.", msg.Name, msg.Password)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	msg := &Message{
		Name:            r.FormValue("name"),
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("confirm_password"),
	}

	if msg.signUpFormValidate() == false {

		//fmt.Printf("%#v\n", msg) // raw Go representation
		//os.Exit(1)

		render(w, "register.html", false, map[string]interface{}{
			"Title": "Home Page",
			"msg":   msg,
		})
		return
	}

	// Process the data (e.g., save to database, send email)
	fmt.Fprintf(w, "Thank you, %s! Your email %s has been received.", msg.Name, msg.Password)

}
