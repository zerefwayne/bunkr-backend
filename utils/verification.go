package utils

import (
	"fmt"
	"log"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/zerefwayne/college-portal-backend/config"
)

func sendMail(email *mail.SGMailV3) (*rest.Response, error) {

	response, err := config.C.SendGrid.Send(email)

	return response, err
}

func getVerificationURL(code string) string {

	url := fmt.Sprintf("%s#/auth/verify/?code=%s", config.C.Env.SendGridEnv.BunkrClientURL, code)

	log.Println(url)

	return url

}

func SendVerificationEmail(code string, name string, email string) (*rest.Response, error) {

	from := mail.NewEmail("Bunkr", "aayushjog@gmail.com")
	subject := "Verify your Bunkr account"
	to := mail.NewEmail(name, email)
	link := fmt.Sprintf("<a href=%q>Click here to verify your email!</a>", getVerificationURL(code))

	message := mail.NewSingleEmail(from, subject, to, "Verify email", link)

	response, err := sendMail(message)

	return response, err

}
