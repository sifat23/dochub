package Controllers

import (
	"dochub/bin/models"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func UserVerify(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	email := r.URL.Query().Get("email")

	fmt.Println("Token:", token, "Email:", email)
	tokenStruct := models.UserToken{
		Token:     token,
		Email:     email,
		ValidTill: nil,
	}

	active := MatchToken(tokenStruct)

	fmt.Println("Active:", active)

}
