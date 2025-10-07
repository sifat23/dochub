package Controllers

import (
	"dochub/bin"
	"dochub/bin/models"
	"dochub/bin/services"
	"dochub/lib"
	"log"
	"net/http"
	"os"
	"time"
)

var pageTitle = "Register Page"

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, "register.html", false, map[string]interface{}{
		"Title": pageTitle,
	})
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	lib.LoadENV() //to load env file

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
		lib.Render(w, "register.html", false, map[string]interface{}{
			"Title": pageTitle,
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

	checkExistToken, err := GetUserToken(user)
	if err != nil {
		log.Println("Error getting user token:", err)
		return
	}

	newToken := services.GenerateToken(user)

	if checkExistToken == nil {
		storeToken, _ := StoreUserToken(*newToken)
		link := os.Getenv("APP_MAIN_URL") + "verification?token=" + storeToken.Token + "&email=" + storeToken.Email
		SendConfirmationMail(map[string]interface{}{
			"Link":    link,
			"Email":   user.Email,
			"Subject": "Account Confirmation Email - DocHub",
		})
	} else {
		deleteToken := DeleteToken(*checkExistToken)
		if deleteToken {
			storeToken, _ := StoreUserToken(*newToken)
			link := os.Getenv("APP_MAIN_URL") + "verification?token=" + storeToken.Token + "&email=" + storeToken.Email
			SendConfirmationMail(map[string]interface{}{
				"Link":    link,
				"Email":   user.Email,
				"Subject": "Account Confirmation Email - DocHub",
			})
		}

	}

	if bin.Db == nil {
		log.Fatal("Database connection is nil! Did you call lib.ConnectToDB() first?")
	}

	// Process the data (e.g., save to database, send email)
	//fmt.Fprintf(w, "Thank you, %s! Your email %s has been received.", msg.Name, msg.Password)

	result, err := bin.Db.Exec("INSERT INTO users (name, email, password, status, created_at) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Status, user.CreatedAt)
	if err != nil {
		lib.Render(w, "register.html", false, map[string]interface{}{
			"Title": pageTitle,
			"error": err,
		})
		return
	}

	id, _ := result.LastInsertId()
	rows, _ := result.RowsAffected()

	lib.Render(w, "register.html", false, map[string]interface{}{
		"Title":   pageTitle,
		"success": "Successfully Register User",
		"id":      id,
		"rows":    rows,
	})
	return
}

func StoreUser() {

}
