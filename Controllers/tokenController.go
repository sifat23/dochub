package Controllers

import (
	"database/sql"
	"dochub/bin"
	"dochub/bin/models"
	"errors"
	"fmt"
	"log"
	"time"
)

func GetUserToken(user models.User) (*models.UserToken, error) {
	var token models.UserToken
	var validTillStr string

	row := bin.Db.QueryRow("SELECT email, token, valid_till FROM user_tokens WHERE email = ?", user.Email)
	err := row.Scan(&token.Email, &token.Token, &validTillStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("sql: no rows found for user %s", user.Email)
			return nil, nil // not an error, but no result
		}
		log.Printf("sql row error: %v", err)
		return nil, err // real error
	}

	t, parseErr := time.Parse("2006-01-02 15:04:05", validTillStr)
	if parseErr != nil {
		log.Printf("time parse error: %v", parseErr)
	} else {
		token.ValidTill = &t
	}

	return &token, nil // success ✅
}

func StoreUserToken(token models.UserToken) (*models.UserToken, error) {
	result, err := bin.Db.Exec("INSERT INTO user_tokens (email, token, valid_till) VALUES (?, ?, ?)", token.Email, token.Token, token.ValidTill)
	if err != nil {
		return nil, err
	}

	rows, _ := result.RowsAffected()
	fmt.Printf("Inserted %d row(s)", rows)

	return &token, nil // ✅ Return token and nil error

}

func MatchToken(token models.UserToken) bool {
	now := time.Now()

	var validTillStr string
	row := bin.Db.QueryRow("SELECT email, token, valid_till FROM user_tokens WHERE email = ? and token = ?", token.Email, token.Token)
	err := row.Scan(&token.Email, &token.Token, &validTillStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("sql: no rows found for user %s", token.Email)
			return false
		}
		log.Printf("sql row error: %v", err)
		return false
	}

	validTill, parseErr := time.Parse("2006-01-02 15:04:05", validTillStr)
	if parseErr != nil {
		log.Printf("time parse error: %v", parseErr)
		return false
	}

	if now.After(validTill) {
		fmt.Println("✅ Token expired at:", validTill)
		return false
	}

	return true
}

func DeleteToken(token models.UserToken) bool {
	del, err := bin.Db.Prepare("DELETE FROM user_tokens WHERE email=?")
	if err != nil {
		panic(err.Error())
		return false
	}
	_, err = del.Exec(token.Email)
	if err != nil {
		return false
	}

	log.Println("DELETE")
	return true
}
