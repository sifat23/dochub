package Controllers

import (
	"bytes"
	"dochub/lib"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"

	"github.com/wneessen/go-mail"
)

func SendConfirmationMail(data interface{}) {
	lib.LoadENV()

	d, ok := data.(map[string]interface{})
	if !ok {
		fmt.Println("Invalid data type — expected map[string]interface{}")
		return
	}

	link := d["Link"].(string)
	email := d["Email"].(string)
	subject := d["Subject"].(string)

	host := os.Getenv("MAIL_HOST")
	from := os.Getenv("MAIL_FROM")
	port := os.Getenv("MAIL_PORT")

	username := os.Getenv("MAIL_ACCOUNT")
	password := os.Getenv("MAIL_PASSWORD")

	tmpl, _ := template.ParseFiles("templates/emails/registration_confirmation.html")
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatalf("failed to execute template: %s", err)
	}

	m := mail.NewMsg()
	if err := m.From(from); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To(email); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextPlain, "Please confirm your account: "+link)
	m.AddAlternativeString(mail.TypeTextHTML, buf.String())

	// Secondly the mail client
	p, _ := strconv.Atoi(port) // convert string to integer
	c, err := mail.NewClient(host,
		mail.WithPort(p), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(username), mail.WithPassword(password))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}

	// Finally let's send out the mail
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}

	fmt.Println("Mail is sent successfully!")
}
