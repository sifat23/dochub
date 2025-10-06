package Controllers

import (
	"dochub/bin"
	"dochub/bin/models"
	"dochub/lib"
	"fmt"
	"log"
	"net/http"
	"time"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, "register.html", false, map[string]interface{}{
		"Title": "Register Page",
	})
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
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
		Name:            r.FormValue("name"),
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("confirm_password"),
	}

	if msg.SignUpFormValidate() == false {

		//fmt.Printf("%#v\n", msg) // raw Go representation
		//os.Exit(1)

		lib.Render(w, "register.html", false, map[string]interface{}{
			"Title": "Home Page",
			"msg":   msg,
		})
		return
	}

	now := time.Now()

	user := models.User{
		Name:      msg.Name,
		Email:     msg.Email,
		Password:  msg.Password,
		Status:    0,
		CreatedAt: &now,
		UpdatedAt: nil,
	}

	if bin.Db == nil {
		log.Fatal("Database connection is nil! Did you call lib.ConnectToDB() first?")
	}

	// Process the data (e.g., save to database, send email)
	//fmt.Fprintf(w, "Thank you, %s! Your email %s has been received.", msg.Name, msg.Password)

	result, err := bin.Db.Exec("INSERT INTO users (name, email, password, status, created_at) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Status, user.CreatedAt)
	if err != nil {
		log.Printf("add User: %v", err)
		return
	}

	id, _ := result.LastInsertId()
	rows, _ := result.RowsAffected()
	fmt.Println("Inserted ID:", id, "Rows:", rows)

	fmt.Printf("%#v\n", "lolo : ", result) // raw Go representation
	//os.Exit(1)
}
