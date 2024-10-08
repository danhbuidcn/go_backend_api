package sendto

import (
	"fmt"
	"log"
	"os"

	"github.com/danhbuidcn/go_backend_api/global"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.uber.org/zap"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "OTP Service"},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp),
	}

	// send using SendGrid
	fromEmail := mail.NewEmail(contentEmail.From.Name, contentEmail.From.Address) // Sender email
	toEmail := mail.NewEmail("Recipient", to[0])                                  // Assume sending to only one recipient

	plainTextContent := fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp)
	htmlContent := fmt.Sprintf("<strong>Your OTP is %s. Please enter it to verify your account.</strong>", otp)

	message := mail.NewSingleEmail(fromEmail, contentEmail.Subject, toEmail, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		global.Logger.Error("Email send failed with SendGrid::", zap.Error(err))
		return err
	}

	if response.StatusCode != 202 {
		// Handle failed response from SendGrid
		log.Printf("Failed to send email, Status Code: %d, Body: %s", response.StatusCode, response.Body)
		return fmt.Errorf("failed to send email, status code: %d", response.StatusCode)
	}

	return nil
}
