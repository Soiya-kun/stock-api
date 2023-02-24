package email_mock

import (
	"fmt"

	"gitlab.com/soy-app/stock-api/usecase/port"
)

type EmailDriver struct{}

func NewEmailDriver() port.Email {
	return &EmailDriver{}
}

func (e EmailDriver) Send(mailAddress []string, subject, body, htmlBody string) error {
	fmt.Println("Email mock send")
	fmt.Println(mailAddress)
	fmt.Println(subject)
	fmt.Println(body)
	fmt.Println(htmlBody)
	return nil
}
