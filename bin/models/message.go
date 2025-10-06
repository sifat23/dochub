package models

import (
	"regexp"
	"strings"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")

type Message struct {
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
	Errors          map[string]string
}

func (msg *Message) SignInFormValidate() bool {
	msg.Errors = make(map[string]string)

	// Email Validation
	match := rxEmail.Match([]byte(msg.Email))
	if match == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}

	// Password Validation
	if strings.TrimSpace(msg.Password) == "" {
		msg.Errors["Password"] = "Please enter password"
	}

	return len(msg.Errors) == 0
}

func (msg *Message) SignUpFormValidate() bool {
	msg.Errors = make(map[string]string)

	// Name Validation
	if strings.TrimSpace(msg.Name) == "" {
		msg.Errors["Name"] = "Please enter your name"
	}

	// Email Validation
	match := rxEmail.Match([]byte(msg.Email))
	if match == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}

	// Password Validation
	if strings.TrimSpace(msg.Password) == "" {
		msg.Errors["Password"] = "Please enter password"
	} else if len(strings.TrimSpace(msg.Password)) < 8 {
		msg.Errors["Password"] = "Password must be at least 8 characters"
	}

	// Confirm Password
	if strings.TrimSpace(msg.Password) != strings.TrimSpace(msg.ConfirmPassword) {
		msg.Errors["Confirm"] = "Password does not match"
	}

	return len(msg.Errors) == 0
}
